import { parseEther } from 'ethers/lib/utils'
import { Config } from '../../boldUpgradeCommon'
import { hoursToBlocks } from './utils'

export const nova: Config = {
  contracts: {
    // it both the excess stake receiver and loser stake escrow
    excessStakeReceiver: '0x40Cd7D713D7ae463f95cE5d342Ea6E7F5cF7C999', // parent to child router
    rollup: '0xFb209827c58283535b744575e11953DCC4bEAD88',
    bridge: '0xC1Ebd02f738644983b6C4B2d440b8e77DdE276Bd',
    sequencerInbox: '0x211E1c4c7f1bF5351Ac850Ed10FD68CFfCF6c21b',
    rollupEventInbox: '0x304807A7ed6c1296df2128E6ff3836e477329CD2',
    outbox: '0xD4B80C3D7240325D18E645B49e6535A3Bf95cc58',
    inbox: '0xc4448b71118c9071Bcb9734A0EAc55D18A153949',
    upgradeExecutor: '0x3ffFbAdAF827559da092217e474760E2b2c3CeDd',
  },
  proxyAdmins: {
    outbox: '0x71d78dc7ccc0e037e12de1e50f5470903ce37148',
    inbox: '0x71d78dc7ccc0e037e12de1e50f5470903ce37148',
    bridge: '0x71d78dc7ccc0e037e12de1e50f5470903ce37148',
    rei: '0x71d78dc7ccc0e037e12de1e50f5470903ce37148',
    seqInbox: '0x71d78dc7ccc0e037e12de1e50f5470903ce37148',
  },
  settings: {
    challengeGracePeriodBlocks: hoursToBlocks(48),
    confirmPeriodBlocks: 45818, // same as old rollup, ~6.4 days
    challengePeriodBlocks: 45818, // same as confirm period
    stakeToken: '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2', // WETH
    // TODO: confirm stakes
    stakeAmt: parseEther('1'),
    miniStakeAmounts: [parseEther('0'), parseEther('1'), parseEther('1')],
    chainId: 42170,
    disableValidatorWhitelist: false,
    blockLeafSize: 2 ** 26, // leaf sizes same as arb1
    bigStepLeafSize: 2 ** 19,
    smallStepLeafSize: 2 ** 23,
    numBigStepLevel: 1,
    maxDataSize: 117964,
    isDelayBufferable: true,
    bufferConfig: {
      max: hoursToBlocks(48), // 2 days
      threshold: hoursToBlocks(1), // well above typical posting frequency
      replenishRateInBasis: 500, // 5% replenishment rate
    },
  },
  // these validators must still be validators on the old rollup during the upgrade, or the upgrade will fail
  validators: [
    // current validators
    '0xE27d4Ed355e5273A3D4855c8e11BC4a8d3e39b87',
    '0x57004b440Cc4eb2FEd8c4d1865FaC907F9150C76',
    '0x24ca61c31c7f9af3ab104db6b9a444f28e9071e3',
  ],
}
