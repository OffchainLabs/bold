import { BigNumber, ethers, Signer, Wallet } from 'ethers'
import fs from 'fs'
import {
  BOLDUpgradeAction__factory,
  Bridge__factory,
  EdgeChallengeManager__factory,
  Outbox__factory,
  RollupAdminLogic__factory,
  RollupEventInbox__factory,
  RollupUserLogic__factory,
  SequencerInbox__factory,
} from '../build/types'
import {
  DeployedContracts,
  Config,
  getJsonFile,
  validateConfig,
} from './common'

const deployDependencies = async (
  signer: Signer
): Promise<
  Omit<DeployedContracts, 'boldAction' | 'preImageHashLookup' | 'rollupReader'>
> => {
  const bridgeFac = new Bridge__factory(signer)
  const bridge = await bridgeFac.deploy()
  console.log(`Bridge implementation deployed at: ${bridge.address}`)

  const seqInboxFac = new SequencerInbox__factory(signer)
  const seqInbox = await seqInboxFac.deploy()
  console.log(`Sequencer inbox implementation deployed at: ${seqInbox.address}`)

  const reiFac = new RollupEventInbox__factory(signer)
  const rei = await reiFac.deploy()
  console.log(`Rollup event inbox implementation deployed at: ${rei.address}`)

  const outboxFac = new Outbox__factory(signer)
  const outbox = await outboxFac.deploy()
  console.log(`Outbox implementation deployed at: ${outbox.address}`)

  // CHRIS: TODO: we need the rollup user logic which was the same as before
  //              except allowed validators to withdraw whilst paused
  const oldRollupUserFac = new RollupUserLogic__factory(signer)
  const oldRollupUser = await oldRollupUserFac.deploy()
  console.log(`Old rollup user logic deployed at: ${oldRollupUser.address}`)

  const newRollupUserFac = new RollupUserLogic__factory(signer)
  const newRollupUser = await newRollupUserFac.deploy()
  console.log(`New rollup user logic deployed at: ${newRollupUser.address}`)

  const newRollupAdminFac = new RollupAdminLogic__factory(signer)
  const newRollupAdmin = await newRollupAdminFac.deploy()
  console.log(`New rollup admin logic deployed at: ${newRollupAdmin.address}`)

  const challengeManagerFac = new EdgeChallengeManager__factory(signer)
  const challengeManager = await challengeManagerFac.deploy()
  console.log(`Challenge manager deployed at: ${challengeManager.address}`)

  return {
    bridge: bridge.address,
    seqInbox: seqInbox.address,
    rei: rei.address,
    outbox: outbox.address,
    oldRollupUser: oldRollupUser.address,
    newRollupUser: newRollupUser.address,
    newRollupAdmin: newRollupAdmin.address,
    challengeManager: challengeManager.address,
  }
}

async function main() {
  const l1RpcVal = process.env.L1_RPC_URL
  if (!l1RpcVal) {
    throw new Error('L1_RPC_URL env variable not set')
  }
  const l1Rpc = new ethers.providers.JsonRpcProvider(l1RpcVal)

  const l1PrivKey = process.env.L1_PRIV_KEY
  if (!l1PrivKey) {
    throw new Error('L1_PRIV_KEY env variable not set')
  }
  const wallet = new Wallet(l1PrivKey, l1Rpc)

  const configLocation = process.env.CONFIG_LOCATION
  if (!configLocation) {
    throw new Error('CONFIG_LOCATION env variable not set')
  }
  const config = getJsonFile(configLocation) as Config
  validateConfig(config)

  const deployedContractsLocation = process.env.DEPLOYED_CONTRACTS_LOCATION
  if (!deployedContractsLocation) {
    throw new Error('DEPLOYED_CONTRACTS_LOCATION env variable not set')
  }

  const deployed = await deployDependencies(wallet)

  const fac = new BOLDUpgradeAction__factory(wallet)
  const boldUpgradeAction = await fac.deploy(
    config.addressReg,
    // CHRIS: TODO: we didnt need any changes in the osp did we? How are we deploying in golang? Lets check that out?
    config.osp,
    // CHRIS: TODO: check these admins are actual contracts
    config.proxyAdmins,
    deployed,
    config.settings
  )
  console.log(`BOLD upgrade action deployed at: ${boldUpgradeAction.address}`)
  const deployedAndBold = {
    ...deployed,
    boldAction: boldUpgradeAction.address,
  }
  console.log(`Deployed contracts written to: ${deployedContractsLocation}`)
  fs.writeFileSync(
    deployedContractsLocation,
    JSON.stringify(deployedAndBold, null, 2)
  )
}

main().then(() => console.log('Done.'))
