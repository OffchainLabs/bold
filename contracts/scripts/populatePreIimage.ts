import { BigNumber, Contract, ethers, Signer, Wallet } from 'ethers'
import fs from 'fs'
import {
  BOLDUpgradeAction__factory,
  Bridge__factory,
  EdgeChallengeManager__factory,
  IOldRollup__factory,
  Outbox__factory,
  RollupAdminLogic__factory,
  RollupEventInbox__factory,
  RollupReader__factory,
  RollupUserLogic__factory,
  SequencerInbox__factory,
  StateHashPreImageLookup__factory,
} from '../build/types'
import {
  DeployedContracts,
  getJsonFile,
  Config,
  validateConfig,
} from './common'
import { Interface } from 'ethers/lib/utils'
import { ExecutionStateStruct } from '../build/types/src/challengeV2/IAssertionChain'

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
  const config = getJsonFile(configLocation)
  validateConfig(config)

  const deployedContractsLocation = process.env.DEPLOYED_CONTRACTS_LOCATION
  if (!deployedContractsLocation) {
    throw new Error('DEPLOYED_CONTRACTS_LOCATION env variable not set')
  }
  const deployedContracts = getJsonFile(
    deployedContractsLocation
  ) as DeployedContracts
  if (!deployedContracts?.preImageHashLookup) {
    throw new Error(
      'preImageHashLookup not found in DEPLOYED_CONTRACTS_LOCATION'
    )
  }

  // find the preimage of the state hash of the latest confirmed node in the rollup
  // and populate it in the lookup
  const iAddressReg = new Interface([
    'function rollup() external returns (address)',
  ])
  const rollupAddr: string = await new Contract(
    config.addressReg,
    iAddressReg,
    wallet
  ).rollup()

  const nodeCreatedEventString =
    'event NodeCreated(uint64 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 indexed nodeHash, bytes32 executionHash, Assertion assertion, bytes32 afterInboxBatchAcc, bytes32 wasmModuleRoot, uint256 inboxMaxCount);'
  const latestConfirmedString =
    'function latestConfirmed() external view returns (uint64)'
  const iOldRollup = new Interface([
    latestConfirmedString,
    nodeCreatedEventString,
  ])
  const latestConfirmed: number = await new Contract(
    rollupAddr,
    iOldRollup,
    wallet
  ).latestConfirmed()

  const latestConfirmedLog = await l1Rpc.getLogs({
    address: rollupAddr,
    fromBlock: 0,
    toBlock: 'latest',
    topics: [
      iOldRollup.getEventTopic('NodeCreated'),
      ethers.utils.hexZeroPad(ethers.utils.hexlify(latestConfirmed), 32),
    ],
  })
  if (latestConfirmedLog.length != 1) {
    throw new Error('Could not find latest confirmed node')
  }
  const latestConfirmedEvent = iOldRollup.parseLog(latestConfirmedLog[0]).args
  const afterState: ExecutionStateStruct =
    latestConfirmedEvent.assertion.afterState
  const inboxCount: BigNumber = latestConfirmedEvent.inboxMaxCount

  const lookup = StateHashPreImageLookup__factory.connect(
    deployedContracts.preImageHashLookup,
    wallet
  )

  const oldRollup = RollupReader__factory.connect(
    deployedContracts.rollupReader,
    l1Rpc
  )
  const node = await oldRollup.getNode(latestConfirmed)
  const stateHash = await lookup.stateHash(afterState, inboxCount)
  if (node.stateHash != stateHash) {
    throw new Error(`State hash mismatch ${node.stateHash} != ${stateHash}}`)
  }

  await lookup.set(stateHash, afterState, inboxCount)
}

// execute this script just prior to execution of the bold upgrade
// it populates the hash lookup contract necessary preimages
main().then(() => console.log('Done.'))
