const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("BHARecharge", function () {
  async function deployFixture() {
    const [owner, user, lpReceiver, burnReceiver, insuranceReceiver] = await ethers.getSigners();
    const signer = ethers.Wallet.createRandom();

    const Token = await ethers.getContractFactory("MockERC20");
    const usdt = await Token.deploy("Tether USD", "USDT", 18);
    const usda = await Token.deploy("Blackhole USD", "USDA", 18);

    const Recharge = await ethers.getContractFactory("BHARecharge");
    const recharge = await Recharge.deploy(signer.address);

    await usdt.mint(user.address, ethers.parseEther("1000"));
    await usda.mint(user.address, ethers.parseEther("1000"));

    return { owner, signer, user, lpReceiver, burnReceiver, insuranceReceiver, usdt, usda, recharge };
  }

  async function signDeposit({ signer, recharge, user, orderNo, tokens, receivers, amounts, expireAt }) {
    const digest = await recharge.depositDigest(user.address, orderNo, tokens, receivers, amounts, expireAt);
    return signer.signingKey.sign(digest).serialized;
  }

  it("distributes a signed multi-token recharge order", async function () {
    const { signer, user, lpReceiver, burnReceiver, insuranceReceiver, usdt, usda, recharge } = await deployFixture();
    const orderNo = "R202607050001";
    const tokens = [await usdt.getAddress(), await usdt.getAddress(), await usda.getAddress()];
    const receivers = [lpReceiver.address, burnReceiver.address, insuranceReceiver.address];
    const amounts = [ethers.parseEther("40"), ethers.parseEther("30"), ethers.parseEther("10")];
    const expireAt = (await ethers.provider.getBlock("latest")).timestamp + 3600;
    const signature = await signDeposit({ signer, recharge, user, orderNo, tokens, receivers, amounts, expireAt });

    await usdt.connect(user).approve(await recharge.getAddress(), ethers.parseEther("70"));
    await usda.connect(user).approve(await recharge.getAddress(), ethers.parseEther("10"));

    await expect(recharge.connect(user).deposit(orderNo, tokens, receivers, amounts, expireAt, signature))
      .to.emit(recharge, "DepositSuccess")
      .withArgs(await recharge.orderHash(orderNo), user.address, orderNo, tokens, receivers, amounts, expireAt);

    expect(await usdt.balanceOf(lpReceiver.address)).to.equal(amounts[0]);
    expect(await usdt.balanceOf(burnReceiver.address)).to.equal(amounts[1]);
    expect(await usda.balanceOf(insuranceReceiver.address)).to.equal(amounts[2]);
  });

  it("rejects changed deposit parameters", async function () {
    const { signer, user, lpReceiver, burnReceiver, usdt, recharge } = await deployFixture();
    const orderNo = "R202607050002";
    const tokens = [await usdt.getAddress(), await usdt.getAddress()];
    const receivers = [lpReceiver.address, burnReceiver.address];
    const amounts = [ethers.parseEther("40"), ethers.parseEther("30")];
    const expireAt = (await ethers.provider.getBlock("latest")).timestamp + 3600;
    const signature = await signDeposit({ signer, recharge, user, orderNo, tokens, receivers, amounts, expireAt });
    const changedAmounts = [ethers.parseEther("41"), ethers.parseEther("30")];

    await usdt.connect(user).approve(await recharge.getAddress(), ethers.parseEther("100"));

    await expect(
      recharge.connect(user).deposit(orderNo, tokens, receivers, changedAmounts, expireAt, signature)
    ).to.be.revertedWithCustomError(recharge, "InvalidSignature");
  });

  it("rejects duplicate order numbers", async function () {
    const { signer, user, lpReceiver, usdt, recharge } = await deployFixture();
    const orderNo = "R202607050003";
    const tokens = [await usdt.getAddress()];
    const receivers = [lpReceiver.address];
    const amounts = [ethers.parseEther("10")];
    const expireAt = (await ethers.provider.getBlock("latest")).timestamp + 3600;
    const signature = await signDeposit({ signer, recharge, user, orderNo, tokens, receivers, amounts, expireAt });

    await usdt.connect(user).approve(await recharge.getAddress(), ethers.parseEther("20"));
    await recharge.connect(user).deposit(orderNo, tokens, receivers, amounts, expireAt, signature);

    await expect(
      recharge.connect(user).deposit(orderNo, tokens, receivers, amounts, expireAt, signature)
    ).to.be.revertedWithCustomError(recharge, "OrderUsed");
  });
});
