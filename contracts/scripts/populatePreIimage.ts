import { ethers, Wallet } from 'ethers'
import { DeployedContracts, getJsonFile, validateConfig } from './common'
import { populateLookup } from './boldUpgradeFunctions'

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

  await populateLookup(
    wallet,
    config.addressReg,
    deployedContracts.preImageHashLookup,
    deployedContracts.rollupReader
  )
}

// execute this script just prior to execution of the bold upgrade
// it populates the hash lookup contract necessary preimages
main().then(() => console.log('Done.'))
