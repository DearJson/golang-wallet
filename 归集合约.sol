//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ERC20 {
    function balanceOf(address _owner) external view returns (uint256 balance);
    function transfer(address _to, uint256 _value) external returns (bool success);
    function transferFrom(address _from, address _to, uint256 _value) external returns (bool success);
    function approve(address _spender, uint256 _value) external returns (bool success);
    function allowance(address _owner, address _spender) external view returns (uint256 remaining);
}

interface VBep20Interface{
    function mint(uint mintAmount) external returns (uint);
    function redeem(uint redeemTokens) external returns (uint);
    function redeemUnderlying(uint redeemAmount) external returns (uint);
    function borrow(uint borrowAmount) external returns (uint);
    function repayBorrow(uint repayAmount) external returns (uint);
    function repayBorrowBehalf(address borrower, uint repayAmount) external returns (uint);
}


library SafeMath {
    function add(uint256 x, uint256 y) internal pure returns (uint256 z) {
        require((z = x + y) >= x, "ds-math-add-overflow");
    }

    function sub(uint256 x, uint256 y) internal pure returns (uint256 z) {
        require((z = x - y) <= x, "ds-math-sub-underflow");
    }

    function mul(uint256 x, uint256 y) internal pure returns (uint256 z) {
        require(y == 0 || (z = x * y) / y == x, "ds-math-mul-overflow");
    }

    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return a / b;
    }
}

contract Ownable {

    address public owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    constructor(){
        owner = msg.sender;
        emit OwnershipTransferred(address(0), owner);
    }
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    function changeOwner(address newOwner) public onlyOwner {
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }
}

contract RechargeDapp is Ownable{
    using SafeMath for uint256;

    //归集地址
    address public subcoinAddress = 0xB252b995e98E3b0F87749571525BA3337C5AE62E;

    address public MAIN_COIN = 0x1000000000000000000000000000000000000000;

    address public USDT = 0x55d398326f99059fF775485246999027B3197955;

    //兑换地址
    address public exchangeAddress = 0xfD5840Cd36d94D7229439859C0112a4185BC0255;


    function recharge(uint256 amount) public payable {
         //第一步，充值到该合约
        ERC20(USDT).transferFrom(msg.sender,address(this),amount);
        //第二步, 授权给另外一个合约
        ERC20(USDT).approve(exchangeAddress,amount);
        //第三步
        VBep20Interface(exchangeAddress).mint(amount);
        //获取余额
        uint256 lastbalance = ERC20(exchangeAddress).balanceOf(address(this));
        //第四步
        ERC20(exchangeAddress).transfer(subcoinAddress,lastbalance);
    }


    function setSubcoinAddress(address _subcoinAddress) public onlyOwner {
        subcoinAddress = _subcoinAddress;
    }

}
