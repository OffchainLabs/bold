import { Contract, ContractReceipt } from 'ethers'
import { ethers } from 'hardhat'
import { Config, DeployedContracts, getConfig, getJsonFile } from './common'
import {
  BOLDUpgradeAction__factory,
  EdgeChallengeManager,
  EdgeChallengeManager__factory,
  RollupUserLogic,
  RollupUserLogic__factory,
} from '../build/types'
import { abi as UpgradeExecutorAbi } from './files/UpgradeExecutor.json'
import dotenv from 'dotenv'
import { RollupMigratedEvent } from '../build/types/src/rollup/BOLDUpgradeAction.sol/BOLDUpgradeAction'
import { abi as OldRollupAbi } from './files/OldRollupUserLogic.json'
import { JsonRpcProvider } from '@ethersproject/providers'
import { getAddress } from 'ethers/lib/utils'

dotenv.config()

type UnwrapPromise<T> = T extends Promise<infer U> ? U : T

type VerificationParams = {
  l1Rpc: JsonRpcProvider
  config: Config
  deployedContracts: DeployedContracts
  preUpgradeState: UnwrapPromise<ReturnType<typeof getPreUpgradeState>>
  receipt: ContractReceipt
}

async function getPreUpgradeState(l1Rpc: JsonRpcProvider, config: Config) {
  const oldRollupContract = new Contract(
    config.contracts.rollup,
    OldRollupAbi,
    l1Rpc
  )

  const stakerCount = await oldRollupContract.stakerCount()

  const stakers: string[] = []
  for (let i = 0; i < stakerCount; i++) {
    stakers.push(await oldRollupContract.getStakerAddress(i))
  }

  return {
    stakers,
  }
}

async function perform(
  l1Rpc: JsonRpcProvider,
  config: Config,
  deployedContracts: DeployedContracts
) {
  await l1Rpc.send('hardhat_impersonateAccount', [
    '0xE6841D92B0C345144506576eC13ECf5103aC7f49'.toLowerCase(),
  ])

  await l1Rpc.send('hardhat_setBalance', [
    '0xE6841D92B0C345144506576eC13ECf5103aC7f49',
    '0x1000000000000000',
  ])

  const timelockImposter = l1Rpc.getSigner(
    '0xE6841D92B0C345144506576eC13ECf5103aC7f49'
  )

  const upExec = new Contract(
    config.contracts.upgradeExecutor,
    UpgradeExecutorAbi,
    timelockImposter
  )
  const boldAction = BOLDUpgradeAction__factory.connect(
    deployedContracts.boldAction,
    timelockImposter
  )

  // what validators did we have in the old rollup?
  const boldActionPerformData = boldAction.interface.encodeFunctionData(
    'perform',
    [config.validators]
  )

  return (await (
    await upExec.execute(deployedContracts.boldAction, boldActionPerformData)
  ).wait()) as ContractReceipt
}

async function verifyPostUpgrade(params: VerificationParams) {
  const { l1Rpc, config, deployedContracts, preUpgradeState, receipt } = params

  const boldAction = BOLDUpgradeAction__factory.connect(
    deployedContracts.boldAction,
    l1Rpc
  )

  const parsedLog = boldAction.interface.parseLog(
    receipt.events![receipt.events!.length - 2]
  ).args as RollupMigratedEvent['args']

  const edgeChallengeManager = EdgeChallengeManager__factory.connect(
    parsedLog.challengeManager,
    l1Rpc
  )

  const newRollup = RollupUserLogic__factory.connect(parsedLog.rollup, l1Rpc)

  await checkBridge(params)
  await checkOldRollup(params)
  await checkNewRollup(params, newRollup)
  await checkNewChallengeManager(params, edgeChallengeManager)
}

async function checkSequencerInbox(params: VerificationParams) {
  const { l1Rpc, config, deployedContracts } = params
  // make sure the impl was updated
  if (
    (await getProxyImpl(l1Rpc, config.contracts.sequencerInbox)) !==
    deployedContracts.seqInbox
  ) {
    throw new Error('SequencerInbox was not upgraded')
  }

  // check delay buffer parameters
  // todo
}

async function checkBridge(params: VerificationParams) {
  const { l1Rpc, config, deployedContracts } = params
  // make sure the impl was updated
  if (
    (await getProxyImpl(l1Rpc, config.contracts.bridge)) !==
    deployedContracts.bridge
  ) {
    throw new Error('Bridge was not upgraded')
  }
}

async function checkOldRollup(params: VerificationParams) {
  const { l1Rpc, config, deployedContracts, preUpgradeState } = params

  const oldRollupContract = new Contract(
    config.contracts.rollup,
    OldRollupAbi,
    l1Rpc
  )

  // ensure the old rollup is paused
  if (!(await oldRollupContract.paused())) {
    throw new Error('Old rollup is not paused')
  }

  // ensure there are no stakers
  if (!(await oldRollupContract.stakerCount()).eq(0)) {
    throw new Error('Old rollup has stakers')
  }

  // ensure that the old stakers are now zombies
  for (const staker of preUpgradeState.stakers) {
    if (!(await oldRollupContract.isZombie(staker))) {
      throw new Error('Old staker is not a zombie')
    }
  }

  // ensure old rollup was upgraded
  if (
    (await getProxyImpl(l1Rpc, config.contracts.rollup, false)) !==
    getAddress(deployedContracts.oldRollupUser)
  ) {
    throw new Error('Old rollup was not upgraded')
  }
}

async function checkNewRollup(
  params: VerificationParams,
  newRollup: RollupUserLogic
) {
  const { l1Rpc, config } = params

  // check stake token address
  if (
    getAddress(await newRollup.stakeToken()) !=
    getAddress(config.settings.stakeToken)
  ) {
    throw new Error('Stake token address does not match')
  }

  // check confirm period blocks
  if (
    !(await newRollup.confirmPeriodBlocks()).eq(
      config.settings.confirmPeriodBlocks
    )
  ) {
    throw new Error('Confirm period blocks does not match')
  }

  // check base stake
  if (!(await newRollup.baseStake()).eq(config.settings.stakeAmt)) {
    throw new Error('Base stake does not match')
  }

  // check fast confirmer
  if (config.settings.anyTrustFastConfirmer.length != 0) {
    if (
      getAddress(await newRollup.anyTrustFastConfirmer()) !==
      getAddress(config.settings.anyTrustFastConfirmer)
    ) {
      throw new Error('Any trust fast confirmer does not match')
    }
  }
}

async function checkNewChallengeManager(
  params: VerificationParams,
  edgeChallengeManager: EdgeChallengeManager
) {
  const { config } = params

  // check stake token address
  if (
    getAddress(await edgeChallengeManager.stakeToken()) !=
    getAddress(config.settings.stakeToken)
  ) {
    throw new Error('Stake token address does not match')
  }

  // check mini stake amounts
  for (let i = 0; i < config.settings.miniStakeAmounts.length; i++) {
    if (
      !(await edgeChallengeManager.stakeAmounts(i)).eq(
        config.settings.miniStakeAmounts[i]
      )
    ) {
      throw new Error('Mini stake amount does not match')
    }
  }

  // check challenge period blocks
  if (
    !(await edgeChallengeManager.challengePeriodBlocks()).eq(
      config.settings.challengePeriodBlocks
    )
  ) {
    throw new Error('Challenge period blocks does not match')
  }

  // check level heights
  if (
    !(await edgeChallengeManager.LAYERZERO_BLOCKEDGE_HEIGHT()).eq(
      config.settings.blockLeafSize
    )
  ) {
    throw new Error('Block leaf size does not match')
  }

  if (
    !(await edgeChallengeManager.LAYERZERO_BIGSTEPEDGE_HEIGHT()).eq(
      config.settings.bigStepLeafSize
    )
  ) {
    throw new Error('Big step leaf size does not match')
  }

  if (
    !(await edgeChallengeManager.LAYERZERO_SMALLSTEPEDGE_HEIGHT()).eq(
      config.settings.smallStepLeafSize
    )
  ) {
    throw new Error('Small step leaf size does not match')
  }

  // check num bigstep levels
  if (
    (await edgeChallengeManager.NUM_BIGSTEP_LEVEL()) !==
    config.settings.numBigStepLevel
  ) {
    throw new Error('Number of big step level does not match')
  }
}

async function getProxyImpl(
  l1Rpc: JsonRpcProvider,
  proxyAddr: string,
  primary = true
) {
  const primarySlot =
    '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'
  const secondarySlot =
    '0x2b1dbce74324248c222f0ec2d5ed7bd323cfc425b336f0253c5ccfda7265546d'
  const val = await l1Rpc.getStorageAt(
    proxyAddr,
    primary ? primarySlot : secondarySlot
  )
  return getAddress('0x' + val.slice(26))
}

async function main() {
  const l1RpcVal = process.env.L1_RPC_URL
  if (!l1RpcVal) {
    throw new Error('L1_RPC_URL env variable not set')
  }
  const l1Rpc = new ethers.providers.JsonRpcProvider(
    l1RpcVal
  ) as JsonRpcProvider

  const deployedContractsLocation = process.env.DEPLOYED_CONTRACTS_LOCATION
  if (!deployedContractsLocation) {
    throw new Error('DEPLOYED_CONTRACTS_LOCATION env variable not set')
  }
  const configLocation = process.env.CONFIG_LOCATION
  if (!configLocation) {
    throw new Error('CONFIG_LOCATION env variable not set')
  }
  const config = await getConfig(configLocation, l1Rpc)

  const deployedContracts = getJsonFile(
    deployedContractsLocation
  ) as DeployedContracts
  if (!deployedContracts.boldAction) {
    throw new Error('No boldAction contract deployed')
  }

  const preUpgradeState = await getPreUpgradeState(l1Rpc, config)
  const receipt = await perform(l1Rpc, config, deployedContracts)
  await verifyPostUpgrade({
    l1Rpc,
    config,
    deployedContracts,
    preUpgradeState,
    receipt,
  })
}

main().then(() => console.log('Done.'))
