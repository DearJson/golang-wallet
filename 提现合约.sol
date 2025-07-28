// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ERC20 {
    function balanceOf(address _owner) external view returns (uint256 balance);
    function transfer(address _to, uint256 _value) external returns (bool success);
    function transferFrom(address _from, address _to, uint256 _value) external returns (bool success);
    function approve(address _spender, uint256 _value) external returns (bool success);
    function allowance(address _owner, address _spender) external view returns (uint256 remaining);
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

//路由合约函数
interface IUniswapV2Router {
    function factory() external view returns (address);

    function getAmountsOut(uint amountIn, address[] calldata path)
    external
    view
    returns (uint[] memory amounts);

    function getAmountsIn(uint amountOut, address[] calldata path)
    external
    view
    returns (uint[] memory amounts);

    function addLiquidity(
        address tokenA,
        address tokenB,
        uint amountADesired,
        uint amountBDesired,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    )
    external
    returns (
        uint amountA,
        uint amountB,
        uint liquidity
    );

    function addLiquidityETH(
        address token,
        uint amountTokenDesired,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    )
    external
    payable
    returns (
        uint amountToken,
        uint amountETH,
        uint liquidity
    );

    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external returns (uint amountA, uint amountB);

    function removeLiquidityETH(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external returns (uint amountToken, uint amountETH);

    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    function swapTokensForExactTokens(
        uint amountOut,
        uint amountInMax,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    function swapExactETHForTokens(
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable returns (uint[] memory amounts);

    function swapTokensForExactETH(
        uint amountOut,
        uint amountInMax,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    function swapExactTokensForETH(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    function swapETHForExactTokens(
        uint amountOut,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable returns (uint[] memory amounts);

    function swapExactTokensForTokensSupportingFeeOnTransferTokens(
        uint amountIn,               // 输入的代币数量
        uint amountOutMin,           // 最小接收到的 USDT 数量（防止滑点）
        address[] calldata path,     // 兑换路径，例如 [代币地址, BNB地址, USDT地址]
        address to,                  // 接收 USDT 的地址
        uint deadline                // 交易的截止时间戳
    ) external;
}

//配对合约
interface IUniswapV2Pair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function swap(
        uint256 amount0Out,
        uint256 amount1Out,
        address to,
        bytes calldata data
    ) external;
}

//工厂合约
interface IUniswapV2Factory {
    function getPair(address token0, address token1) external returns (address);
}



contract Ownable {
    address public owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    constructor(){
        owner = msg.sender;
        emit OwnershipTransferred(address(0), owner);
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Ownable: caller is not the owner");
        _;
    }

    function changeOwner(address newOwner) public onlyOwner {
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }
}

contract WithdrawDapp is Ownable {
    using SafeMath for uint256;

    address public USDT = 0x55d398326f99059fF775485246999027B3197955;//USDT合约地址

    address public TOKEN = 0xE26a492cED210bD84e11147e8AeC84B23390f22e;//SPA合约地址

    address public PANCAKE_SWAP_ROUTER = 0x10ED43C718714eb63d5aA57B78B54704E256024E;//固定薄饼路由合约地址

    mapping(string => string) private abiMap;

    address public uniswapV2Pair;//配对合约地址2
        constructor() {
        // 初始化时添加多种 ABI
        abiMap["withdraw1"] = '[{"internalType": "uint256", "name": "amount", "type": "uint256"},{"internalType": "address", "name": "withdrawAddress", "type": "address"}]';
        abiMap["withdraw2"] = '[{"internalType": "uint256", "name": "amount", "type": "uint256"},{"internalType": "address", "name": "withdrawAddress", "type": "address"}]';
    }

    //充值50%的释放 ，提现的时候 100%买币进入对应的提现地址
    //商家收到积分后，闪兑成USDT提现， 50%USDT到提现钱包， 50%买币到提现钱包
    function withdraw(string memory abiKey, bytes memory data) public payable {
        if (keccak256(abi.encodePacked(abiKey)) == keccak256(abi.encodePacked("withdraw1"))){
            (uint256 amount,address withdrawAddress) = abi.decode(data, (uint256,address));
            ERC20(USDT).transferFrom(msg.sender, address(this), amount);
            ERC20(USDT).approve(PANCAKE_SWAP_ROUTER, amount);
            address[] memory buyPath = new address[](2);
            buyPath[0] = USDT;
            buyPath[1] = TOKEN;
            IUniswapV2Router(PANCAKE_SWAP_ROUTER).swapExactTokensForTokensSupportingFeeOnTransferTokens(amount,0,buyPath,withdrawAddress,block.timestamp);
        }else if(keccak256(abi.encodePacked(abiKey)) == keccak256(abi.encodePacked("withdraw2"))){
            (uint256 amount,address withdrawAddress) = abi.decode(data, (uint256,address));
            ERC20(USDT).transferFrom(msg.sender, address(this), amount);
       
            uint256 amount1 = amount.mul(50).div(100);
            if (amount1 > 0){
                address[] memory buyPath = new address[](2);
                buyPath[0] = USDT;
                buyPath[1] = TOKEN;
                IUniswapV2Router(PANCAKE_SWAP_ROUTER).swapExactTokensForTokensSupportingFeeOnTransferTokens(amount1,0,buyPath,withdrawAddress,block.timestamp);
            }
            uint256 amount2 =  amount.sub(amount1);
            if (amount2 > 0){
                ERC20(USDT).transferFrom(msg.sender, withdrawAddress, amount);
            }
        }else{
            revert("Operator Error");
        }
    }

}
