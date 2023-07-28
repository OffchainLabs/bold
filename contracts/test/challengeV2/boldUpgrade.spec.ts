// import { Address, getL2Network } from "@arbitrum/sdk";
// import { ArbitrumProvider } from "@arbitrum/sdk/dist/lib/utils/arbProvider";
import { JsonRpcProvider, Provider } from '@ethersproject/providers'
import { expect } from 'chai'
import { Contract, ContractFactory, Signer, Wallet, constants } from 'ethers'
import { HDNode, arrayify, parseEther } from 'ethers/lib/utils'
import {
  ProxyAdmin__factory,
  TransparentUpgradeableProxy__factory,
} from '../../build/types'
import { getJsonFile } from '../../scripts/common'
import fs from 'fs'
import path from 'path'
import {
  ERC1967Proxy__factory,
  RollupAdminLogic__factory,
  RollupProxy__factory,
  RollupUserLogic__factory,
} from '../../build/types'
import {
  abi as UpgradeExecutorAbi,
  bytecode as UpgradeExecutorBytecode,
} from '../../scripts/files/UpgradeExecutor.json'
import { deployBoldUpgrade } from '../../scripts/boldUpgradeFunctions'

const wait = async (ms: number) => new Promise(res => setTimeout(res, ms))

const mineBlock = async (signer: Signer) => {
  await (
    await signer.sendTransaction({ to: await signer.getAddress(), value: 0 })
  ).wait()
}

const localL2Rpc = 'http://localhost:8549'
const localL1Rpc = 'http://localhost:8545'

const l1mnemonic =
  'indoor dish desk flag debris potato excuse depart ticket judge file exit'

const transferOrGetToUpgradeExec = async (
  rollupAdmin: Wallet,
  rollupAddress: string
) => {
  const rollupOwner = await RollupUserLogic__factory.connect(
    rollupAddress,
    rollupAdmin
  ).owner()
  if (rollupOwner.toLowerCase() !== rollupAdmin.address.toLowerCase()) {
    const upExec = new Contract(rollupOwner, UpgradeExecutorAbi, rollupAdmin)
    const execRole = (await upExec.EXECUTOR_ROLE()) as string
    const hasRole = (await upExec.hasRole(
      execRole,
      rollupAdmin.address
    )) as boolean
    if (hasRole != true)
      throw Error('Upgrade executor does not have EXECUTOR role')
    return upExec
  } else {
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

    console.log(upExec.address)
    console.log(
      await RollupUserLogic__factory.connect(
        rollupAddress,
        rollupAdmin
      ).owner(),
      rollupAdmin.address
    )
    await RollupAdminLogic__factory.connect(
      rollupAddress,
      rollupAdmin
    ).setOwner(upExec.address)

    return upExec
  }
}

// set this test up to execute on ci - we're gonna need to mostly test there i think
// just be careful with each change/addition
// CHRIS: TODO: check all comments in these test

describe.only('BoldUpgradeTest', () => {
  it('Can upgrade', async () => {
    // we can only run this test once!!!, oh dear. We want to restart the chain if that's the case?

    const localNetworksPath = path.join(
      __dirname,
      '../../scripts/files/localNetwork.json'
    )
    const localNetworks = await getJsonFile(localNetworksPath)

    // we need something else deployed? nope
    // we need to deploy an upgrade executor? yes
    // do we have the priv key to the rollup owner?
    // i'll assume yes
    const l1Rpc = new JsonRpcProvider(localL1Rpc)

    const rollupAdmin = new Wallet(
      HDNode.fromMnemonic(l1mnemonic).derivePath("m/44'/60'/0'/0/1").privateKey
    ).connect(l1Rpc)
    console.log(rollupAdmin.privateKey)
    return;

    const rollupAddr = localNetworks['l2Network']['ethBridge']['rollup']
    console.log(
      await RollupUserLogic__factory.connect(rollupAddr, l1Rpc).owner()
    )
    const upExec = await transferOrGetToUpgradeExec(rollupAdmin, rollupAddr)

    // deploy and address reg, what's in that? - eek, can we do this without the address reg
    // yes?

      // deployBoldUpgrade(
      //   rollupAdmin, {

      //   }
      // )

    

    

    // we have an existing rollup available
    // we want to upgrade it by running the bold upgrade action
    // aftwards we check that something actually happened

    // how can we run the action? Do we have an upgrade executor owning the chain?
    // if not, we need to set one up and make it the owner of the l1 contracts
  })
})
