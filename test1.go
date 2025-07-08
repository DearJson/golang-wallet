package main

import (
	"fmt"
	"gfast/library"
	"github.com/shopspring/decimal"
	"strings"
)

func main() {
	amountString := "0x" + strings.TrimLeft(library.SubStr("0xa9059cbb0000000000000000000000001b86a64098daac00842e0c8d1cc053978a502cd70000000000000000000000000000000000000000000000030461545f7e995f09", -64, -1), "0")

	//amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).Div(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals)))).RoundFloor(int32(coinAddress[contractAddress].Decimals)).String()
	//dapp-sus
	//amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, 18)), 18).Mul(decimal.NewFromFloat32(0.97)).Truncate(18).String()
	amount := decimal.NewFromBigInt(library.HexToBigInt(amountString), int32(-(18))).Truncate(18).String()

	fmt.Printf("%s", amount)
}
