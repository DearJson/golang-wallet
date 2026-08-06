// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IWithdrawERC20 {
    function transfer(address to, uint256 value) external returns (bool);
}

/// @title BHAWithdraw
/// @notice Custody withdrawal contract executed by an approved backend job.
/// @dev The project funds this contract first. After off-chain approval, the
/// owner job calls withdraw with a unique order number. The contract prevents
/// duplicate orders and emits events for backend reconciliation.
contract BHAWithdraw {
    address public owner;

    mapping(bytes32 => bool) public usedOrderHashes;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event WithdrawSuccess(
        bytes32 indexed orderHash,
        address indexed operator,
        string orderNo,
        address indexed token,
        address to,
        uint256 amount
    );
    event RescueToken(address indexed token, address indexed to, uint256 amount);

    error NotOwner();
    error InvalidOrder();
    error InvalidAddress();
    error InvalidAmount();
    error OrderUsed();
    error TokenTransferFailed(address token);

    modifier onlyOwner() {
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    constructor() {
        owner = msg.sender;
        emit OwnershipTransferred(address(0), msg.sender);
    }

    /// @notice Execute one approved withdrawal order.
    function withdraw(
        string calldata orderNo,
        address token,
        address to,
        uint256 amount
    ) external onlyOwner {
        bytes32 computedOrderHash = _orderHash(orderNo);
        if (token == address(0) || to == address(0)) revert InvalidAddress();
        if (amount == 0) revert InvalidAmount();
        if (usedOrderHashes[computedOrderHash]) revert OrderUsed();

        _safeTransfer(token, to, amount);
        usedOrderHashes[computedOrderHash] = true;

        emit WithdrawSuccess(computedOrderHash, msg.sender, orderNo, token, to, amount);
    }

    function orderHash(string calldata orderNo) external pure returns (bytes32) {
        return _orderHash(orderNo);
    }

    function transferOwnership(address newOwner) external onlyOwner {
        if (newOwner == address(0)) revert InvalidAddress();
        address previousOwner = owner;
        owner = newOwner;
        emit OwnershipTransferred(previousOwner, newOwner);
    }

    /// @notice Rescue tokens that should not remain in this contract.
    function rescueToken(address token, address to, uint256 amount) external onlyOwner {
        if (token == address(0) || to == address(0)) revert InvalidAddress();
        if (amount == 0) revert InvalidAmount();
        _safeTransfer(token, to, amount);
        emit RescueToken(token, to, amount);
    }

    function _orderHash(string calldata orderNo) internal pure returns (bytes32) {
        if (bytes(orderNo).length == 0) revert InvalidOrder();
        return keccak256(bytes(orderNo));
    }

    function _safeTransfer(address token, address to, uint256 amount) internal {
        (bool success, bytes memory data) = token.call(
            abi.encodeWithSelector(IWithdrawERC20.transfer.selector, to, amount)
        );
        if (!success || (data.length != 0 && !abi.decode(data, (bool)))) {
            revert TokenTransferFailed(token);
        }
    }
}
