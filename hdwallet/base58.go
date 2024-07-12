package hdwallet

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/shengdoushi/base58"
)

var tronAlphabet = base58.NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Encode(input []byte) string {
	return base58.Encode(input, tronAlphabet)
}

func EncodeCheck(input []byte) string {
	h256h0 := sha256.New()
	h256h0.Write(input)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	inputCheck := input
	inputCheck = append(inputCheck, h1[:4]...)

	return Encode(inputCheck)
}

func Decode(input string) ([]byte, error) {
	return base58.Decode(input, tronAlphabet)
}

func DecodeCheck(input string) ([]byte, error) {
	decodeCheck, err := Decode(input)

	if err != nil {
		return nil, err
	}

	if len(decodeCheck) < 4 {
		return nil, fmt.Errorf("addres base58 not check ok")
	}

	decodeData := decodeCheck[:len(decodeCheck)-4]

	h256h0 := sha256.New()
	h256h0.Write(decodeData)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	if h1[0] == decodeCheck[len(decodeData)] &&
		h1[1] == decodeCheck[len(decodeData)+1] &&
		h1[2] == decodeCheck[len(decodeData)+2] &&
		h1[3] == decodeCheck[len(decodeData)+3] {
		return decodeData, nil
	}

	return nil, fmt.Errorf("addres hash not check ok")
}

// FromHexAddress 41 ---- > T
func FromHexAddress(hexAddress string) (string, error) {
	addrByte, err := hex.DecodeString(hexAddress)
	if err != nil {
		return "", err
	}

	sha := sha256.New()
	sha.Write(addrByte)
	shaStr := sha.Sum(nil)

	sha2 := sha256.New()
	sha2.Write(shaStr)
	shaStr2 := sha2.Sum(nil)

	addrByte = append(addrByte, shaStr2[:4]...)

	return string(base58Encode(addrByte)), nil
}

var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// base58Encode 编码
func base58Encode(input []byte) []byte {
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}
	var result []byte
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabets[mod.Int64()])
	}
	reverseBytes(result)
	return result
}

// base58Decode 解码
func base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte(base58Alphabets, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == base58Alphabets[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded[:len(decoded)-4]
}

// reverseBytes 翻转字节
func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
