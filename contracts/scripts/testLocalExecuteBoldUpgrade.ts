import { Contract, ContractFactory, Wallet, ethers } from 'ethers'
import { DeployedContracts, getJsonFile } from './common'
import fs from 'fs'
import path from 'path'
import {
  ProxyAdmin__factory,
  TransparentUpgradeableProxy__factory,
  RollupAdminLogic__factory,
  BOLDUpgradeAction__factory,
  Bridge__factory,
} from '../build/types'
import {
  abi as UpgradeExecutorAbi,
  bytecode as UpgradeExecutorBytecode,
} from './files/UpgradeExecutor.json'
import dotenv from 'dotenv'
import { HDNode } from 'ethers/lib/utils'

dotenv.config()

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

  const deployedContractsLocation = process.env.DEPLOYED_CONTRACTS_LOCATION
  if (!deployedContractsLocation) {
    throw new Error('DEPLOYED_CONTRACTS_LOCATION env variable not set')
  }

  const deployedContracts = getJsonFile(
    deployedContractsLocation
  ) as DeployedContracts
  if (!deployedContracts.boldAction) {
    throw new Error('No boldAction contract deployed')
  }
  if (!deployedContracts.upgradeExecutor) {
    throw new Error('No upgradeExecutor contract deployed')
  }

  const upExec = new Contract(
    deployedContracts.upgradeExecutor,
    UpgradeExecutorAbi,
    wallet
  )
  const boldAction = BOLDUpgradeAction__factory.connect(
    deployedContracts.boldAction,
    wallet
  )
  const boldActionPerformData =
    boldAction.interface.encodeFunctionData('perform')
  const rollup = await boldAction.ROLLUP()
  const oldRollup = RollupAdminLogic__factory.connect(rollup, wallet)

  await (
    await upExec.execute(deployedContracts.boldAction, boldActionPerformData)
  ).wait()
}

main().then(() => console.log('Done.'))
