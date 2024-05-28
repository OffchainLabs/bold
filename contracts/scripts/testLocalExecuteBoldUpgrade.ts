import { Contract, ContractReceipt } from 'ethers'
import { ethers as ethers2 } from 'hardhat'
import { DeployedContracts, getConfig, getJsonFile } from './common'
import fs from 'fs'
import { BOLDUpgradeAction__factory } from '../build/types'
import { abi as UpgradeExecutorAbi } from './files/UpgradeExecutor.json'
import dotenv from 'dotenv'
import { RollupMigratedEvent } from '../build/types/src/rollup/BOLDUpgradeAction.sol/BOLDUpgradeAction'

dotenv.config()

async function main() {
  const l1RpcVal = process.env.L1_RPC_URL
  if (!l1RpcVal) {
    throw new Error('L1_RPC_URL env variable not set')
  }
  const l1Rpc = new ethers2.providers.JsonRpcProvider(l1RpcVal)


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

  await l1Rpc.send(
    "hardhat_impersonateAccount",
    ["0xE6841D92B0C345144506576eC13ECf5103aC7f49".toLowerCase()],
  )

  await l1Rpc.send(
    "hardhat_setBalance",
    ["0xE6841D92B0C345144506576eC13ECf5103aC7f49", '0x1000000000000000'],
  )

  const timelockImposter = l1Rpc.getSigner('0xE6841D92B0C345144506576eC13ECf5103aC7f49'.toLowerCase())
  
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

  const receipt = (await (
    await upExec.execute(deployedContracts.boldAction, boldActionPerformData)
  ).wait()) as ContractReceipt

  const parsedLog = boldAction.interface.parseLog(
    receipt.events![receipt.events!.length - 2]
  ).args as RollupMigratedEvent['args']

  console.log(`Deployed contracts written to: ${deployedContractsLocation}`)
  fs.writeFileSync(
    deployedContractsLocation,
    JSON.stringify(
      {
        ...deployedContracts,
        newEdgeChallengeManager: parsedLog.challengeManager,
      },
      null,
      2
    )
  )
}

main().then(() => console.log('Done.'))
