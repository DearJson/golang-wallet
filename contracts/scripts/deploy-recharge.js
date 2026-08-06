const hre = require("hardhat");

async function main() {
  const signer = process.env.RECHARGE_SIGNER;
  if (!signer) {
    throw new Error("Missing RECHARGE_SIGNER");
  }

  const Recharge = await hre.ethers.getContractFactory("BHARecharge");
  const recharge = await Recharge.deploy(signer);
  await recharge.waitForDeployment();

  console.log(`BHARecharge deployed to: ${await recharge.getAddress()}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
