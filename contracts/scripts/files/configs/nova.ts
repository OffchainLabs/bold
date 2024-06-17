import { parseEther } from 'ethers/lib/utils'
import { Config } from '../../common'
import { hoursToBlocks } from '.'

export const nova: Config = {
  contracts: {
    // the l1Timelock does not actually need to be the timelock
    // it is only used to set the excess stake receiver / loser stake escrow
    // TODO: change this to a fee router before real deployment
    l1Timelock: '0xE6841D92B0C345144506576eC13ECf5103aC7f49',
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
    stakeToken: '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
    stakeAmt: parseEther('1'), // warrants discussion
    miniStakeAmounts: [
      parseEther('0'),
      parseEther('1'),
      parseEther('1'),
    ],
    chainId: 42161,
    anyTrustFastConfirmer: '0x0000000000000000000000000000000000000000', // TODO
    disableValidatorWhitelist: false,
    blockLeafSize: 2**26, // leaf sizes same as arb1
    bigStepLeafSize: 2**19,
    smallStepLeafSize: 2**23,
    numBigStepLevel: 1,
    maxDataSize: 117964,
    isDelayBufferable: true,
    bufferConfig: {
      max: hoursToBlocks(48),
      threshold: hoursToBlocks(1),
      replenishRateInBasis: 500,
    },
  },
  validators: [ // current validators
    '0xB51EDdfc9A945e2B909905e4F242C4796Ac0C61d',
    '0x54c0D3d6C101580dB3be8763A2aE2c6bb9dc840c',
    '0x658e8123722462F888b6fa01a7dbcEFe1D6DD709',
    '0xDfB23DFE9De7dcC974467195C8B7D5cd21C9d7cB'
  ],
}
