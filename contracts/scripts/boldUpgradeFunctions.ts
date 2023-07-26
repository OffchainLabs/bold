// import { BigNumber, Contract, ContractFactory, ethers, Signer } from 'ethers'
// import {
//   BOLDUpgradeAction__factory,
//   Bridge__factory,
//   EdgeChallengeManager__factory,
//   Outbox__factory,
//   RollupAdminLogic__factory,
//   RollupEventInbox__factory,
//   RollupReader__factory,
//   RollupUserLogic__factory,
//   SequencerInbox__factory,
//   StateHashPreImageLookup__factory,
// } from '../build/types'
// import { DeployedContracts, Config } from './common'
// import { Interface } from 'ethers/lib/utils'
// import { ExecutionStateStruct } from '../build/types/src/challengeV2/IAssertionChain'
// import {
//   abi as OldRollupAbi,
//   bytecode as OldRollupBytecode,
// } from './files/OldRollupUserLogic.json'

// export const deployDependencies = async (
//   signer: Signer,
//   log: boolean = false
// ): Promise<
//   Omit<DeployedContracts, 'boldAction' | 'preImageHashLookup' | 'rollupReader'>
// > => {
//   const bridgeFac = new Bridge__factory(signer)
//   const bridge = await bridgeFac.deploy()
//   if (log) {
//     console.log(`Bridge implementation deployed at: ${bridge.address}`)
//   }

//   const seqInboxFac = new SequencerInbox__factory(signer)
//   const seqInbox = await seqInboxFac.deploy()
//   if (log) {
//     console.log(
//       `Sequencer inbox implementation deployed at: ${seqInbox.address}`
//     )
//   }

//   const reiFac = new RollupEventInbox__factory(signer)
//   const rei = await reiFac.deploy()
//   if (log) {
//     console.log(`Rollup event inbox implementation deployed at: ${rei.address}`)
//   }

//   const outboxFac = new Outbox__factory(signer)
//   const outbox = await outboxFac.deploy()
//   if (log) {
//     console.log(`Outbox implementation deployed at: ${outbox.address}`)
//   }

//   const oldRollupUserFac = new ContractFactory(
//     OldRollupAbi,
//     OldRollupBytecode,
//     signer
//   )
//   const oldRollupUser = await oldRollupUserFac.deploy()
//   if (log) {
//     console.log(`Old rollup user logic deployed at: ${oldRollupUser.address}`)
//   }

//   const newRollupUserFac = new RollupUserLogic__factory(signer)
//   const newRollupUser = await newRollupUserFac.deploy()
//   if (log) {
//     console.log(`New rollup user logic deployed at: ${newRollupUser.address}`)
//   }

//   const newRollupAdminFac = new RollupAdminLogic__factory(signer)
//   const newRollupAdmin = await newRollupAdminFac.deploy()
//   if (log) {
//     console.log(`New rollup admin logic deployed at: ${newRollupAdmin.address}`)
//   }

//   const challengeManagerFac = new EdgeChallengeManager__factory(signer)
//   const challengeManager = await challengeManagerFac.deploy()
//   if (log) {
//     console.log(`Challenge manager deployed at: ${challengeManager.address}`)
//   }

//   return {
//     bridge: bridge.address,
//     seqInbox: seqInbox.address,
//     rei: rei.address,
//     outbox: outbox.address,
//     oldRollupUser: oldRollupUser.address,
//     newRollupUser: newRollupUser.address,
//     newRollupAdmin: newRollupAdmin.address,
//     challengeManager: challengeManager.address,
//   }
// }

// export const deployBoldUpgrade = async (
//   wallet: Signer,
//   config: Config,
//   log: boolean = false
// ): Promise<DeployedContracts> => {
//   const deployed = await deployDependencies(wallet, log)

//   const fac = new BOLDUpgradeAction__factory(wallet)
//   const boldUpgradeAction = await fac.deploy(
//     config.addressReg,
//     // CHRIS: TODO: we didnt need any changes in the osp did we? How are we deploying in golang? Lets check that out?
//     config.osp,
//     config.proxyAdmins,
//     deployed,
//     config.settings
//   )
//   if (log) {
//     console.log(`BOLD upgrade action deployed at: ${boldUpgradeAction.address}`)
//   }
//   const deployedAndBold = {
//     ...deployed,
//     boldAction: boldUpgradeAction.address,
//     rollupReader: await boldUpgradeAction.ROLLUP_READER(),
//     preImageHashLookup: await boldUpgradeAction.PREIMAGE_LOOKUP(),
//   }

//   return deployedAndBold
// }

// export const populateLookup = async (
//   wallet: Signer,
//   addressRegAddr: string,
//   preImageHashLookupAddr: string,
//   rollupReaderAddr: string
// ) => {
//   // find the preimage of the state hash of the latest confirmed node in the rollup
//   // and populate it in the lookup
//   const iAddressReg = new Interface([
//     'function rollup() external returns (address)',
//   ])
//   const rollupAddr: string = await new Contract(
//     addressRegAddr,
//     iAddressReg,
//     wallet
//   ).rollup()

//   const nodeCreatedEventString =
//     'event NodeCreated(uint64 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 indexed nodeHash, bytes32 executionHash, Assertion assertion, bytes32 afterInboxBatchAcc, bytes32 wasmModuleRoot, uint256 inboxMaxCount);'
//   const latestConfirmedString =
//     'function latestConfirmed() external view returns (uint64)'
//   const iOldRollup = new Interface([
//     latestConfirmedString,
//     nodeCreatedEventString,
//   ])
//   const latestConfirmed: number = await new Contract(
//     rollupAddr,
//     iOldRollup,
//     wallet
//   ).latestConfirmed()

//   const latestConfirmedLog = await wallet.provider!.getLogs({
//     address: rollupAddr,
//     fromBlock: 0,
//     toBlock: 'latest',
//     topics: [
//       iOldRollup.getEventTopic('NodeCreated'),
//       ethers.utils.hexZeroPad(ethers.utils.hexlify(latestConfirmed), 32),
//     ],
//   })
//   if (latestConfirmedLog.length != 1) {
//     throw new Error('Could not find latest confirmed node')
//   }
//   const latestConfirmedEvent = iOldRollup.parseLog(latestConfirmedLog[0]).args
//   const afterState: ExecutionStateStruct =
//     latestConfirmedEvent.assertion.afterState
//   const inboxCount: BigNumber = latestConfirmedEvent.inboxMaxCount

//   const lookup = StateHashPreImageLookup__factory.connect(
//     preImageHashLookupAddr,
//     wallet
//   )

//   const oldRollup = RollupReader__factory.connect(rollupReaderAddr, wallet)
//   const node = await oldRollup.getNode(latestConfirmed)
//   const stateHash = await lookup.stateHash(afterState, inboxCount)
//   if (node.stateHash != stateHash) {
//     throw new Error(`State hash mismatch ${node.stateHash} != ${stateHash}}`)
//   }

//   await lookup.set(stateHash, afterState, inboxCount)
// }
