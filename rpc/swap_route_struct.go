package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

//根据精确token交换尽量多的token
type swapExactTokensForTokens struct {
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	To           common.Address
	Deadline     *big.Int
}

//使用尽量少的token交换精确的token
type swapTokensForExactTokens struct {
	AmountOut   *big.Int
	AmountInMax *big.Int
	Path        []common.Address
	To          common.Address
	Deadline    *big.Int
}

//根据精确的token交换尽量多的token
type swapExactETHForTokens struct {
	AmountOutMin *big.Int
	Path         []common.Address
	To           common.Address
	Deadline     *big.Int
}

//根据尽量少的token交换精确的ETH
type swapTokensForExactETH struct {
	AmountOut   *big.Int
	AmountInMax *big.Int
	Path        []common.Address
	To          common.Address
	Deadline    *big.Int
}

//根据精确的token交换尽量多的ETH
type swapExactTokensForETH struct {
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	To           common.Address
	Deadline     *big.Int
}

//根据精确的ETH交换尽量多的TOKEN
type swapETHForExactTokens struct {
	AmountOut *big.Int
	Path      []common.Address
	To        common.Address
	Deadline  *big.Int
}
