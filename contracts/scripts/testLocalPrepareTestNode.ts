import { JsonRpcProvider, Provider } from '@ethersproject/providers'
import { expect } from 'chai'
import {
  Contract,
  ContractFactory,
  Signer,
  Wallet,
  constants,
  ethers,
} from 'ethers'
import { HDNode, arrayify, parseEther } from 'ethers/lib/utils'
import { getJsonFile } from './common'
import fs from 'fs'
import path from 'path'
import {
  ProxyAdmin__factory,
  TransparentUpgradeableProxy__factory,
  ERC1967Proxy__factory,
  RollupAdminLogic__factory,
  RollupProxy__factory,
  RollupUserLogic__factory,
} from '../build/types'
import {
  abi as UpgradeExecutorAbi,
  bytecode as UpgradeExecutorBytecode,
} from './files/UpgradeExecutor.json'
import { deployBoldUpgrade } from './boldUpgradeFunctions'
import dotenv from "dotenv";

dotenv.config();

const wait = async (ms: number) => new Promise(res => setTimeout(res, ms))

const mineBlock = async (signer: Signer) => {
  await (
    await signer.sendTransaction({ to: await signer.getAddress(), value: 0 })
  ).wait()
}

const transferToUpgradeExec = async (
  rollupAdmin: Wallet,
  rollupAddress: string
) => {
  const upgradeExecutorImpl = await new ContractFactory(
    UpgradeExecutorAbi,
    UpgradeExecutorBytecode,
    rollupAdmin
  ).deploy()
  await upgradeExecutorImpl.deployed()

  const proxyAdmin = await new ProxyAdmin__factory(rollupAdmin).deploy()
  await proxyAdmin.deployed()

  const upExecProxy = await new TransparentUpgradeableProxy__factory(
    rollupAdmin
  ).deploy(upgradeExecutorImpl.address, proxyAdmin.address, '0x')
  await upExecProxy.deployed()

  const upExec = new Contract(
    upExecProxy.address,
    UpgradeExecutorAbi,
    rollupAdmin
  )
  await upExec.initialize(rollupAdmin.address, [rollupAdmin.address])

  await RollupAdminLogic__factory.connect(rollupAddress, rollupAdmin).setOwner(
    upExec.address
  )

  return upExec
}

// set this test up to execute on ci - we're gonna need to mostly test there i think
// just be careful with each change/addition
// CHRIS: TODO: check all comments in these test

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

  // are we creating the config? no, i think we can populate it
  // we want to transfer to the upgrade exec
  //   // load the local network
  const localNetworksPath = path.join(
    __dirname,
    './files/localNetwork.json'
  )
  const localNetworks = await getJsonFile(localNetworksPath)
  const rollupAddr = localNetworks['l2Network']['ethBridge']['rollup']
  const upExec = await transferToUpgradeExec(wallet, rollupAddr);

//   const configLocation = process.env.CONFIG_LOCATION
//   if (!configLocation) {
//     throw new Error('CONFIG_LOCATION env variable not set')
//   }
//   const config = getJsonFile(configLocation) as Config
//   validateConfig(config)

//   const deployedContractsLocation = process.env.DEPLOYED_CONTRACTS_LOCATION
//   if (!deployedContractsLocation) {
//     throw new Error('DEPLOYED_CONTRACTS_LOCATION env variable not set')
//   }

//   const deployedAndBold = await deployBoldUpgrade(wallet, config, true)

//   console.log(`Deployed contracts written to: ${deployedContractsLocation}`)
//   fs.writeFileSync(
//     deployedContractsLocation,
//     JSON.stringify(deployedAndBold, null, 2)
//   )
}

main().then(() => console.log('Done.'))
