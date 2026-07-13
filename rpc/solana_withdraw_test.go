package rpc

import (
	"crypto/ed25519"
	"testing"

	"gfast/hdwallet"
)

func TestNewPreparedSolanaTransactionUsesSerializedSignature(t *testing.T) {
	seed := make([]byte, ed25519.SeedSize)
	for i := range seed {
		seed[i] = byte(i + 1)
	}
	privateKey := ed25519.NewKeyFromSeed(seed)
	message := []byte("withdraw-order-123")
	signedTx := signSolanaTransaction(message, privateKey)

	prepared, err := newPreparedSolanaTransaction(signedTx, 12345)
	if err != nil {
		t.Fatalf("newPreparedSolanaTransaction() error = %v", err)
	}
	wantSignature := hdwallet.SolanaBase58Encode(ed25519.Sign(privateKey, message))
	if prepared.Signature != wantSignature {
		t.Fatalf("Signature = %q, want %q", prepared.Signature, wantSignature)
	}
	if prepared.LastValidBlockHeight != 12345 {
		t.Fatalf("LastValidBlockHeight = %d, want 12345", prepared.LastValidBlockHeight)
	}
}

func TestNewPreparedSolanaTransactionRejectsInvalidData(t *testing.T) {
	if _, err := newPreparedSolanaTransaction([]byte{0}, 0); err == nil {
		t.Fatal("newPreparedSolanaTransaction() error = nil, want error")
	}
}
