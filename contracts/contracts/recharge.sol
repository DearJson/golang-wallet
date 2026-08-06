// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IERC20 {
    function transferFrom(address from, address to, uint256 value) external returns (bool);
    function transfer(address to, uint256 value) external returns (bool);
}

/// @title BHARecharge
/// @notice Signed multi-token, multi-recipient recharge entrypoint.
/// @dev Backend creates an order and signs the exact deposit parameters. The
/// contract verifies the signature, distributes tokens, marks the order used,
/// and emits a log that the backend can poll by orderHash/orderNo.
contract BHARecharge {
    uint256 private constant SECP256K1_N_DIV_2 =
        0x7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0;

    address public owner;
    address public signer;
    bool public paused;

    mapping(bytes32 => bool) public usedOrderHashes;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event SignerUpdated(address indexed previousSigner, address indexed newSigner);
    event PausedUpdated(bool paused);
    event DepositSuccess(
        bytes32 indexed orderHash,
        address indexed user,
        string orderNo,
        address[] tokens,
        address[] receivers,
        uint256[] amounts,
        uint256 expireAt
    );
    event RescueToken(address indexed token, address indexed to, uint256 amount);

    error NotOwner();
    error Paused();
    error InvalidSigner();
    error InvalidOrder();
    error InvalidArrayLength();
    error InvalidAddress();
    error InvalidAmount();
    error OrderExpired();
    error OrderUsed();
    error InvalidSignature();
    error TokenTransferFailed(address token);

    modifier onlyOwner() {
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    modifier whenNotPaused() {
        if (paused) revert Paused();
        _;
    }

    constructor(address initialSigner) {
        if (initialSigner == address(0)) revert InvalidSigner();
        owner = msg.sender;
        signer = initialSigner;
        emit OwnershipTransferred(address(0), msg.sender);
        emit SignerUpdated(address(0), initialSigner);
    }

    /// @notice Deposit one order and distribute ERC20 tokens to all receivers.
    /// @param orderNo Backend order number.
    /// @param tokens Token address for each distribution line. Repeating tokens
    /// is supported, which lets one order split the same token to many receivers.
    /// @param receivers Receiver address for each distribution line.
    /// @param amounts Token amount for each distribution line, in token base units.
    /// @param expireAt Unix timestamp after which the signed order is invalid.
    /// @param signature Backend signer signature over the exact parameters.
    function deposit(
        string calldata orderNo,
        address[] calldata tokens,
        address[] calldata receivers,
        uint256[] calldata amounts,
        uint256 expireAt,
        bytes calldata signature
    ) external whenNotPaused {
        bytes32 computedOrderHash = _orderHash(orderNo);
        _validateDeposit(computedOrderHash, tokens, receivers, amounts, expireAt);

        bytes32 digest = depositDigest(msg.sender, orderNo, tokens, receivers, amounts, expireAt);
        if (_recoverSigner(digest, signature) != signer) revert InvalidSignature();

        usedOrderHashes[computedOrderHash] = true;

        for (uint256 i = 0; i < tokens.length; i++) {
            _safeTransferFrom(tokens[i], msg.sender, receivers[i], amounts[i]);
        }

        emit DepositSuccess(computedOrderHash, msg.sender, orderNo, tokens, receivers, amounts, expireAt);
    }

    /// @notice Returns the digest the backend must sign.
    function depositDigest(
        address user,
        string calldata orderNo,
        address[] calldata tokens,
        address[] calldata receivers,
        uint256[] calldata amounts,
        uint256 expireAt
    ) public view returns (bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(
                block.chainid,
                address(this),
                user,
                keccak256(bytes(orderNo)),
                keccak256(abi.encode(tokens)),
                keccak256(abi.encode(receivers)),
                keccak256(abi.encode(amounts)),
                expireAt
            )
        );
        return _toEthSignedMessageHash(structHash);
    }

    function orderHash(string calldata orderNo) external pure returns (bytes32) {
        return _orderHash(orderNo);
    }

    function setSigner(address newSigner) external onlyOwner {
        if (newSigner == address(0)) revert InvalidSigner();
        address previousSigner = signer;
        signer = newSigner;
        emit SignerUpdated(previousSigner, newSigner);
    }

    function setPaused(bool value) external onlyOwner {
        paused = value;
        emit PausedUpdated(value);
    }

    function transferOwnership(address newOwner) external onlyOwner {
        if (newOwner == address(0)) revert InvalidAddress();
        address previousOwner = owner;
        owner = newOwner;
        emit OwnershipTransferred(previousOwner, newOwner);
    }

    /// @notice Rescue tokens sent to this contract by mistake.
    function rescueToken(address token, address to, uint256 amount) external onlyOwner {
        if (token == address(0) || to == address(0)) revert InvalidAddress();
        _safeTransfer(token, to, amount);
        emit RescueToken(token, to, amount);
    }

    function _validateDeposit(
        bytes32 orderHash_,
        address[] calldata tokens,
        address[] calldata receivers,
        uint256[] calldata amounts,
        uint256 expireAt
    ) internal view {
        if (orderHash_ == bytes32(0)) revert InvalidOrder();
        if (tokens.length == 0) revert InvalidArrayLength();
        if (tokens.length != receivers.length || tokens.length != amounts.length) revert InvalidArrayLength();
        if (block.timestamp > expireAt) revert OrderExpired();
        if (usedOrderHashes[orderHash_]) revert OrderUsed();

        for (uint256 i = 0; i < tokens.length; i++) {
            if (tokens[i] == address(0) || receivers[i] == address(0)) revert InvalidAddress();
            if (amounts[i] == 0) revert InvalidAmount();
        }
    }

    function _orderHash(string calldata orderNo) internal pure returns (bytes32) {
        if (bytes(orderNo).length == 0) revert InvalidOrder();
        return keccak256(bytes(orderNo));
    }

    function _recoverSigner(bytes32 digest, bytes calldata signature) internal pure returns (address) {
        if (signature.length != 65) revert InvalidSignature();

        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := calldataload(signature.offset)
            s := calldataload(add(signature.offset, 0x20))
            v := byte(0, calldataload(add(signature.offset, 0x40)))
        }

        if (v < 27) {
            v += 27;
        }
        if (v != 27 && v != 28) revert InvalidSignature();
        if (uint256(s) > SECP256K1_N_DIV_2) revert InvalidSignature();

        address recovered = ecrecover(digest, v, r, s);
        if (recovered == address(0)) revert InvalidSignature();
        return recovered;
    }

    function _toEthSignedMessageHash(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }

    function _safeTransferFrom(address token, address from, address to, uint256 amount) internal {
        (bool success, bytes memory data) = token.call(
            abi.encodeWithSelector(IERC20.transferFrom.selector, from, to, amount)
        );
        if (!success || (data.length != 0 && !abi.decode(data, (bool)))) {
            revert TokenTransferFailed(token);
        }
    }

    function _safeTransfer(address token, address to, uint256 amount) internal {
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(IERC20.transfer.selector, to, amount));
        if (!success || (data.length != 0 && !abi.decode(data, (bool)))) {
            revert TokenTransferFailed(token);
        }
    }
}
