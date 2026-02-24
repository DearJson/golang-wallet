package hdwallet

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/shengdoushi/base58"
)

var solanaAlphabet = base58.NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// GenerateSolanaKeyPair 生成Solana Ed25519密钥对
// 返回: 私钥hex字符串, Solana地址(Base58), error
func GenerateSolanaKeyPair() (privateKeyHex string, address string, err error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", fmt.Errorf("generate ed25519 key error: %v", err)
	}
	// 私钥存储为hex字符串（64字节 = 128 hex字符）
	privateKeyHex = hex.EncodeToString(priv)
	// Solana地址 = 公钥的Base58编码
	address = base58.Encode(pub, solanaAlphabet)
	return
}

// SolanaAddressFromPrivateKeyHex 从私钥hex字符串推导Solana地址
func SolanaAddressFromPrivateKeyHex(privateKeyHex string) (string, error) {
	privBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("decode private key hex error: %v", err)
	}
	if len(privBytes) != ed25519.PrivateKeySize {
		return "", fmt.Errorf("invalid private key length: %d", len(privBytes))
	}
	priv := ed25519.PrivateKey(privBytes)
	pub := priv.Public().(ed25519.PublicKey)
	return base58.Encode(pub, solanaAlphabet), nil
}

// SolanaBase58Decode 解码Solana Base58字符串为字节
func SolanaBase58Decode(input string) ([]byte, error) {
	return base58.Decode(input, solanaAlphabet)
}

// SolanaBase58Encode 编码字节为Solana Base58字符串
func SolanaBase58Encode(input []byte) string {
	return base58.Encode(input, solanaAlphabet)
}

// GetSolanaPrivateKey 从Base58或hex字符串获取Ed25519私钥
func GetSolanaPrivateKey(privateKey string) (ed25519.PrivateKey, error) {
	var privBytes []byte
	var err error

	// 先尝试Base58解码（Solana最常用格式）
	privBytes, err = SolanaBase58Decode(privateKey)
	if err != nil {
		// Base58失败，尝试hex解码
		privBytes, err = hex.DecodeString(privateKey)
		if err != nil {
			return nil, fmt.Errorf("decode private key error (tried base58 and hex): %v", err)
		}
	}

	if len(privBytes) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("invalid private key length: %d, expected: %d", len(privBytes), ed25519.PrivateKeySize)
	}
	return ed25519.PrivateKey(privBytes), nil
}
