import { parseEther } from 'ethers/lib/utils'
import { Config } from '../../common'

export const local: Config = {
  contracts: {
    bridge: '0x5eCF728ffC5C5E802091875f96281B5aeECf6C49',
    inbox: '0x9f8c1c641336A371031499e3c362e40d58d0f254',
    outbox: '0x50143333b44Ea46255BEb67255C9Afd35551072F',
    rollup: '0xC3124dD1FA0e5D6135c25279760DBF9d9286467B',
    sequencerInbox: '0x18d19C5d3E685f5be5b9C86E097f0E439285D216',
    rollupEventInbox: '0x0e73faf857e1ca53e700856fcf19f31f920a1e3c',
    upgradeExecutor: '0x513d9f96d4d0563debae8a0dc307ea0e46b10ed7',
    l1Timelock: '0xC3124dD1FA0e5D6135c25279760DBF9d9286467B',
  },
  proxyAdmins: {
    outbox: '0x2a1f38c9097e7883570e0b02bfbe6869cc25d8a3',
    inbox: '0x2a1f38c9097e7883570e0b02bfbe6869cc25d8a3',
    bridge: '0x2a1f38c9097e7883570e0b02bfbe6869cc25d8a3',
    rei: '0x2a1f38c9097e7883570e0b02bfbe6869cc25d8a3',
    seqInbox: '0x2a1f38c9097e7883570e0b02bfbe6869cc25d8a3',
  },
  settings: {
    challengeGracePeriodBlocks: 10,
    confirmPeriodBlocks: 100,
    challengePeriodBlocks: 110,
    stakeToken: '0x408Da76E87511429485C32E4Ad647DD14823Fdc4',
    stakeAmt: parseEther('1'),
    miniStakeAmounts: [
      parseEther('6'),
      parseEther('5'),
      parseEther('4'),
      parseEther('3'),
      parseEther('2'),
      parseEther('1'),
    ],
    chainId: 412346,
    anyTrustFastConfirmer: '0x6d903f6003cca6255D85CcA4D3B5E5146dC33925',
    disableValidatorWhitelist: true,
    blockLeafSize: 1048576,
    bigStepLeafSize: 512,
    smallStepLeafSize: 128,
    numBigStepLevel: 4,
    maxDataSize: 117964,
    isDelayBufferable: true,
    bufferConfig: {
      max: 14400,
      threshold: 300,
      replenishRateInBasis: 500,
    },
  },
  validators: ['0xf10EF80c6eF4930A62C5F9661c91339Df4dBB173'],
}
