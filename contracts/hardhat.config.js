require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config({ quiet: true });

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  defaultNetwork: "bsc",
  solidity: {
    version: "0.8.20",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  },
  networks: {
    bsc: {
      url: process.env.BSC_RPC_URL || "",
      chainId: 56,
      accounts: process.env.DEPLOY_PRIVATE_KEY ? [process.env.DEPLOY_PRIVATE_KEY] : []
    }
  }
};
