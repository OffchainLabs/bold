// import { Address, getL2Network } from "@arbitrum/sdk";
// import { ArbitrumProvider } from "@arbitrum/sdk/dist/lib/utils/arbProvider";
import { JsonRpcProvider, Provider } from "@ethersproject/providers";
import { expect } from "chai";
import { Signer, Wallet, constants } from "ethers";
import { parseEther } from "ethers/lib/utils";
// import {
//   ProxyAdmin__factory,
// } from "../../build/types";
import { getJsonFile } from "../../scripts/common";

const wait = async (ms: number) => new Promise((res) => setTimeout(res, ms));

const mineBlock = async (signer: Signer) => {
  await (await signer.sendTransaction({ to: await signer.getAddress(), value: 0 })).wait();
};

describe.only("BoldUpgradeTest", () => {

    it("Can upgrade", async () => {
        const jsFile = await getJsonFile("../files/localNework.json");
        console.log(jsFile);
        // we have an existing rollup available
        // we want to upgrade it by running the bold upgrade action
        // aftwards we check that something actually happened

        // how can we run the action? Do we have an upgrade executor owning the chain?
        // if not, we need to set one up and make it the owner of the l1 contracts


        





        

    })

})

