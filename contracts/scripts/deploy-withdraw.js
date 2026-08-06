const hre = require("hardhat");

async function main() {
  const Withdraw = await hre.ethers.getContractFactory("BHAWithdraw");
  const withdraw = await Withdraw.deploy();
  await withdraw.waitForDeployment();

  console.log(`BHAWithdraw deployed to: ${await withdraw.getAddress()}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
