const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("BHAWithdraw", function () {
  async function deployFixture() {
    const [owner, user, attacker] = await ethers.getSigners();

    const Token = await ethers.getContractFactory("MockERC20");
    const usdt = await Token.deploy("Tether USD", "USDT", 18);

    const Withdraw = await ethers.getContractFactory("BHAWithdraw");
    const withdraw = await Withdraw.deploy();

    await usdt.mint(owner.address, ethers.parseEther("1000"));
    await usdt.transfer(await withdraw.getAddress(), ethers.parseEther("1000"));

    return { owner, user, attacker, usdt, withdraw };
  }

  it("lets owner withdraw funded token to an approved user", async function () {
    const { owner, user, usdt, withdraw } = await deployFixture();
    const orderNo = "W202607050001";
    const amount = ethers.parseEther("88");
    const token = await usdt.getAddress();

    await expect(withdraw.connect(owner).withdraw(orderNo, token, user.address, amount))
      .to.emit(withdraw, "WithdrawSuccess")
      .withArgs(await withdraw.orderHash(orderNo), owner.address, orderNo, token, user.address, amount);

    expect(await usdt.balanceOf(user.address)).to.equal(amount);
  });

  it("rejects non-owner withdraw calls", async function () {
    const { attacker, user, usdt, withdraw } = await deployFixture();

    await expect(
      withdraw.connect(attacker).withdraw("W202607050002", await usdt.getAddress(), user.address, ethers.parseEther("1"))
    ).to.be.revertedWithCustomError(withdraw, "NotOwner");
  });

  it("rejects duplicate withdraw order numbers", async function () {
    const { owner, user, usdt, withdraw } = await deployFixture();
    const orderNo = "W202607050003";
    const token = await usdt.getAddress();

    await withdraw.connect(owner).withdraw(orderNo, token, user.address, ethers.parseEther("1"));

    await expect(
      withdraw.connect(owner).withdraw(orderNo, token, user.address, ethers.parseEther("1"))
    ).to.be.revertedWithCustomError(withdraw, "OrderUsed");
  });

  it("does not mark order as used when token transfer fails", async function () {
    const { owner, user, usdt, withdraw } = await deployFixture();
    const orderNo = "W202607050004";
    const orderHash = await withdraw.orderHash(orderNo);

    await expect(
      withdraw.connect(owner).withdraw(orderNo, await usdt.getAddress(), user.address, ethers.parseEther("1001"))
    ).to.be.revertedWithCustomError(withdraw, "TokenTransferFailed");

    expect(await withdraw.usedOrderHashes(orderHash)).to.equal(false);
  });
});
