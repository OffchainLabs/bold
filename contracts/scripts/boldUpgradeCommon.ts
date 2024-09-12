import { BigNumber, providers } from 'ethers'
import { parseEther } from 'ethers/lib/utils'
import fs from 'fs'

import { configs } from './files/configs'

export interface DeployedContracts {
  bridge: string
  seqInbox: string
  rei: string
  outbox: string
  inbox: string
  oldRollupUser: string
  newRollupUser: string
  newRollupAdmin: string
  challengeManager: string
  boldAction: string
  rollupReader: string
  preImageHashLookup: string
  prover0: string
  proverMem: string
  proverMath: string
  proverHostIo: string
  osp: string
}

export const getJsonFile = (fileLocation: string) => {
  return JSON.parse(fs.readFileSync(fileLocation).toString())
}

export const getConfig = async (
  configName: string,
  l1Rpc: providers.Provider
): Promise<Config> => {
  const config = configs[configName as keyof typeof configs]
  if (!config) {
    throw new Error('config not found')
  }
  await validateConfig(config, l1Rpc)
  return config
}

export interface Config {
  contracts: {
    excessStakeReceiver: string
    rollup: string
    bridge: string
    sequencerInbox: string
    rollupEventInbox: string
    outbox: string
    inbox: string
    upgradeExecutor: string
  }
  proxyAdmins: {
    outbox: string
    inbox: string
    bridge: string
    rei: string
    seqInbox: string
  }
  settings: {
    challengeGracePeriodBlocks: number
    confirmPeriodBlocks: number
    challengePeriodBlocks: number
    stakeToken: string
    stakeAmt: BigNumber
    miniStakeAmounts: BigNumber[]
    chainId: number
    disableValidatorWhitelist: boolean
    maxDataSize: number
    blockLeafSize: number
    bigStepLeafSize: number
    smallStepLeafSize: number
    numBigStepLevel: number
    isDelayBufferable: boolean
    bufferConfig: {
      max: number
      threshold: number
      replenishRateInBasis: number
    }
  }
  validators: string[]
}

export type RawConfig = Omit<Config, 'settings'> & {
  settings: Omit<Config['settings'], 'stakeAmt' | 'miniStakeAmounts'> & {
    stakeAmt: string
    miniStakeAmounts: string[]
  }
}

export const validateConfig = async (
  config: Config,
  l1Rpc: providers.Provider
) => {
  // check all the config.contracts exist
  if ((await l1Rpc.getCode(config.contracts.excessStakeReceiver)).length <= 2) {
    throw new Error('excessStakeReceiver address is not a contract')
  }
  if ((await l1Rpc.getCode(config.contracts.rollup)).length <= 2) {
    throw new Error('rollup address is not a contract')
  }
  if ((await l1Rpc.getCode(config.contracts.bridge)).length <= 2) {
    throw new Error('bridge address is not a contract')
  }
  if ((await l1Rpc.getCode(config.contracts.sequencerInbox)).length <= 2) {
    throw new Error('sequencerInbox address is not a contract')
  }
  if ((await l1Rpc.getCode(config.contracts.rollupEventInbox)).length <= 2) {
    throw new Error('rollupEventInbox address is not a contract')
  }
  if ((await l1Rpc.getCode(config.contracts.outbox)).length <= 2) {
    throw new Error('outbox address is not a contract')
  }
  if ((await l1Rpc.getCode(config.contracts.inbox)).length <= 2) {
    throw new Error('inbox address is not a contract')
  }
  if ((await l1Rpc.getCode(config.contracts.upgradeExecutor)).length <= 2) {
    throw new Error('upgradeExecutor address is not a contract')
  }

  // check all the config.proxyAdmins exist
  if ((await l1Rpc.getCode(config.proxyAdmins.outbox)).length <= 2) {
    throw new Error('outbox proxy admin address is not a contract')
  }
  if ((await l1Rpc.getCode(config.proxyAdmins.bridge)).length <= 2) {
    throw new Error('bridge proxy admin address is not a contract')
  }
  if ((await l1Rpc.getCode(config.proxyAdmins.rei)).length <= 2) {
    throw new Error('rei proxy admin address is not a contract')
  }
  if ((await l1Rpc.getCode(config.proxyAdmins.seqInbox)).length <= 2) {
    throw new Error('seqInbox proxy admin address is not a contract')
  }

  // check all the settings exist
  if (config.settings.confirmPeriodBlocks === 0) {
    throw new Error('confirmPeriodBlocks is 0')
  }
  if (config.settings.stakeToken.length === 0) {
    throw new Error('stakeToken address is empty')
  }
  if (config.settings.chainId === 0) {
    throw new Error('chainId is 0')
  }
  if (config.settings.blockLeafSize === 0) {
    throw new Error('blockLeafSize is 0')
  }
  if (config.settings.bigStepLeafSize === 0) {
    throw new Error('bigStepLeafSize is 0')
  }
  if (config.settings.smallStepLeafSize === 0) {
    throw new Error('smallStepLeafSize is 0')
  }
  if (config.settings.numBigStepLevel === 0) {
    throw new Error('numBigStepLevel is 0')
  }

  const stakeAmount = BigNumber.from(config.settings.stakeAmt)
  // check it's more than 1 eth
  if (stakeAmount.lt(parseEther('1'))) {
    throw new Error('stakeAmt is less than 1 eth')
  }
  const miniStakeAmounts = config.settings.miniStakeAmounts.map(BigNumber.from)

  if (miniStakeAmounts.length !== config.settings.numBigStepLevel + 2) {
    throw new Error('miniStakeAmts length is not numBigStepLevel + 2')
  }

  if (
    !config.settings.disableValidatorWhitelist &&
    config.validators.length === 0
  ) {
    throw new Error('no validators')
  }
}
