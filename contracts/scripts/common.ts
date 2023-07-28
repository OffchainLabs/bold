import { BigNumber } from 'ethers'
import fs from 'fs'
export interface DeployedContracts {
  bridge: string
  seqInbox: string
  rei: string
  outbox: string
  oldRollupUser: string
  newRollupUser: string
  newRollupAdmin: string
  challengeManager: string
  boldAction: string
  rollupReader: string
  preImageHashLookup: string
  upgradeExecutor?: string
  newEdgeChallengeManager?: string
}

export const getJsonFile = (fileLocation: string) => {
  return JSON.parse(fs.readFileSync(fileLocation).toString())
}

export interface Config {
  contracts: {
    l1Timelock: string
    rollup: string
    bridge: string
    sequencerInbox: string
    rollupEventInbox: string
    outbox: string
    inbox: string
    osp: string
  }
  proxyAdmins: {
    outbox: string
    bridge: string
    rei: string
    seqInbox: string
  }
  settings: {
    confirmPeriodBlocks: number
    stakeToken: string
    stakeAmt: BigNumber
    miniStakeAmt: BigNumber
    chainId: number
    anyTrustFastConfirmer: string
  }
}

export const validateConfig = (config: Config) => {
  // CHRIS: TODO: ensure vals are valid
  // CHRIS: TODO: check the proxy admins are actual contracts
}
