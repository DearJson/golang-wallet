package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"strings"
	"testing"

	"gfast/hdwallet"
	"gfast/library"
)

func TestDecryptSolanaPrivateKey(t *testing.T) {
	seed := make([]byte, ed25519.SeedSize)
	for i := range seed {
		seed[i] = byte(i + 1)
	}
	privateKey := hex.EncodeToString(ed25519.NewKeyFromSeed(seed))
	encrypted, err := library.EncryptByAes([]byte(privateKey))
	if err != nil {
		t.Fatalf("EncryptByAes() error = %v", err)
	}

	var output bytes.Buffer
	if err = decryptSolanaPrivateKey(strings.NewReader(encrypted+"\n"), &output); err != nil {
		t.Fatalf("decryptSolanaPrivateKey() error = %v", err)
	}
	if got := strings.TrimSpace(output.String()); got != privateKey {
		t.Fatalf("decrypted private key = %q, want %q", got, privateKey)
	}
}

func TestDecryptSolanaBase58PrivateKey(t *testing.T) {
	seed := make([]byte, ed25519.SeedSize)
	for i := range seed {
		seed[i] = byte(i + 1)
	}
	privateKey := hdwallet.SolanaBase58Encode(ed25519.NewKeyFromSeed(seed))
	encrypted, err := library.EncryptByAes([]byte(privateKey))
	if err != nil {
		t.Fatalf("EncryptByAes() error = %v", err)
	}

	var output bytes.Buffer
	if err = decryptSolanaPrivateKey(strings.NewReader(encrypted), &output); err != nil {
		t.Fatalf("decryptSolanaPrivateKey() error = %v", err)
	}
	if got := strings.TrimSpace(output.String()); got != privateKey {
		t.Fatalf("decrypted private key = %q, want %q", got, privateKey)
	}
}

func TestDecryptSolanaPrivateKeyRejectsInvalidCiphertext(t *testing.T) {
	var output bytes.Buffer
	if err := decryptSolanaPrivateKey(strings.NewReader("not-base64"), &output); err == nil {
		t.Fatal("decryptSolanaPrivateKey() error = nil, want error")
	}
}
