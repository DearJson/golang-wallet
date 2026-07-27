package rpc

import (
	"crypto/ed25519"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"testing"

	"gfast/hdwallet"
)

func newSolanaWithdrawTestKey() ed25519.PrivateKey {
	seed := make([]byte, ed25519.SeedSize)
	for i := range seed {
		seed[i] = byte(i + 1)
	}
	return ed25519.NewKeyFromSeed(seed)
}

func solanaTestAddress(fill byte) string {
	key := make([]byte, ed25519.PublicKeySize)
	for i := range key {
		key[i] = fill
	}
	return hdwallet.SolanaBase58Encode(key)
}

func parseTestTransferInstructions(t *testing.T, rawTransaction []byte, wantAccountCount int, wantReadonlyUnsigned byte) []solanaInstruction {
	t.Helper()
	if len(rawTransaction) < 68 || rawTransaction[0] != 1 {
		t.Fatalf("invalid serialized transaction: %x", rawTransaction)
	}
	message := rawTransaction[65:]
	if message[0] != 1 || message[1] != 0 || message[2] != wantReadonlyUnsigned {
		t.Fatalf("message header = %v, want [1 0 %d]", message[:3], wantReadonlyUnsigned)
	}
	offset := 3
	accountCount := int(message[offset])
	offset++
	if accountCount != wantAccountCount {
		t.Fatalf("account count = %d, want %d", accountCount, wantAccountCount)
	}
	offset += accountCount*ed25519.PublicKeySize + 32
	instructionCount := int(message[offset])
	offset++
	instructions := make([]solanaInstruction, 0, instructionCount)
	for i := 0; i < instructionCount; i++ {
		programIDIndex := message[offset]
		offset++
		accountLen := int(message[offset])
		offset++
		accounts := append([]byte(nil), message[offset:offset+accountLen]...)
		offset += accountLen
		dataLen := int(message[offset])
		offset++
		data := append([]byte(nil), message[offset:offset+dataLen]...)
		offset += dataLen
		instructions = append(instructions, solanaInstruction{programIDIndex: programIDIndex, accounts: accounts, data: data})
	}
	return instructions
}

func TestNewPreparedSolanaTransactionUsesSerializedSignature(t *testing.T) {
	privateKey := newSolanaWithdrawTestKey()
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

func TestPrepareSOLTransfersBuildsTwoOutputs(t *testing.T) {
	blockhash := solanaTestAddress(9)
	client := &SolanaRpcClient{rpcCallOverride: func(method string, _ interface{}) (json.RawMessage, error) {
		switch method {
		case "getBalance":
			return json.RawMessage(`{"context":{"slot":1},"value":1000000000}`), nil
		case "getLatestBlockhash":
			return json.RawMessage(fmt.Sprintf(`{"context":{"slot":1},"value":{"blockhash":%q,"lastValidBlockHeight":123}}`, blockhash)), nil
		default:
			return nil, fmt.Errorf("unexpected RPC method %q", method)
		}
	}}

	privateKey := newSolanaWithdrawTestKey()
	prepared, err := client.PrepareSOLTransfers(hdwallet.SolanaBase58Encode(privateKey), []SolanaTransfer{
		{ToAddress: solanaTestAddress(2), Amount: 98_000_000},
		{ToAddress: solanaTestAddress(3), Amount: 2_000_000},
	})
	if err != nil {
		t.Fatalf("PrepareSOLTransfers() error = %v", err)
	}
	instructions := parseTestTransferInstructions(t, prepared.RawTransaction, 4, 1)
	if len(instructions) != 2 {
		t.Fatalf("instruction count = %d, want 2", len(instructions))
	}
	for i, wantAmount := range []uint64{98_000_000, 2_000_000} {
		if instructions[i].programIDIndex != 3 || len(instructions[i].data) != 12 || binary.LittleEndian.Uint32(instructions[i].data[:4]) != 2 {
			t.Fatalf("instruction %d = %#v, want system transfer", i, instructions[i])
		}
		if got := binary.LittleEndian.Uint64(instructions[i].data[4:]); got != wantAmount {
			t.Fatalf("instruction %d amount = %d, want %d", i, got, wantAmount)
		}
	}
}

func TestPrepareSPLTokenTransfersBuildsTwoOutputs(t *testing.T) {
	privateKey := newSolanaWithdrawTestKey()
	fromAddress := hdwallet.SolanaBase58Encode(privateKey.Public().(ed25519.PublicKey))
	mint := solanaTestAddress(8)
	blockhash := solanaTestAddress(9)
	destinationOwners := []string{solanaTestAddress(2), solanaTestAddress(3)}
	accountInfoCall := 0
	client := &SolanaRpcClient{rpcCallOverride: func(method string, _ interface{}) (json.RawMessage, error) {
		switch method {
		case "getAccountInfo":
			owner := fromAddress
			if accountInfoCall > 0 {
				owner = destinationOwners[accountInfoCall-1]
			}
			accountInfoCall++
			return json.RawMessage(fmt.Sprintf(`{"value":{"data":{"parsed":{"info":{"mint":%q,"owner":%q,"tokenAmount":{"amount":"1000000000","decimals":6}},"type":"account"},"program":"spl-token"},"owner":%q}}`, mint, owner, TokenProgramIDBase58)), nil
		case "getBalance":
			return json.RawMessage(`{"context":{"slot":1},"value":1000000000}`), nil
		case "getLatestBlockhash":
			return json.RawMessage(fmt.Sprintf(`{"context":{"slot":1},"value":{"blockhash":%q,"lastValidBlockHeight":123}}`, blockhash)), nil
		default:
			return nil, fmt.Errorf("unexpected RPC method %q", method)
		}
	}}

	prepared, err := client.PrepareSPLTokenTransfers(hdwallet.SolanaBase58Encode(privateKey), mint, []SolanaTransfer{
		{ToAddress: destinationOwners[0], Amount: 98_000_000},
		{ToAddress: destinationOwners[1], Amount: 2_000_000},
	})
	if err != nil {
		t.Fatalf("PrepareSPLTokenTransfers() error = %v", err)
	}
	instructions := parseTestTransferInstructions(t, prepared.RawTransaction, 5, 1)
	if len(instructions) != 2 {
		t.Fatalf("instruction count = %d, want 2", len(instructions))
	}
	for i, wantAmount := range []uint64{98_000_000, 2_000_000} {
		if instructions[i].programIDIndex != 4 || len(instructions[i].data) != 9 || instructions[i].data[0] != 3 {
			t.Fatalf("instruction %d = %#v, want SPL transfer", i, instructions[i])
		}
		if got := binary.LittleEndian.Uint64(instructions[i].data[1:]); got != wantAmount {
			t.Fatalf("instruction %d amount = %d, want %d", i, got, wantAmount)
		}
	}
}

func TestPrepareSPLTokenTransfersCreatesMissingDestinationATA(t *testing.T) {
	privateKey := newSolanaWithdrawTestKey()
	fromAddress := hdwallet.SolanaBase58Encode(privateKey.Public().(ed25519.PublicKey))
	mint := solanaTestAddress(8)
	blockhash := solanaTestAddress(9)
	destinationOwners := []string{solanaTestAddress(2), solanaTestAddress(3)}
	accountInfoCall := 0
	client := &SolanaRpcClient{rpcCallOverride: func(method string, _ interface{}) (json.RawMessage, error) {
		switch method {
		case "getAccountInfo":
			if accountInfoCall == 2 {
				accountInfoCall++
				return json.RawMessage(`{"value":null}`), nil
			}
			owner := fromAddress
			if accountInfoCall == 1 {
				owner = destinationOwners[0]
			}
			accountInfoCall++
			return json.RawMessage(fmt.Sprintf(`{"value":{"data":{"parsed":{"info":{"mint":%q,"owner":%q,"tokenAmount":{"amount":"1000000000","decimals":6}},"type":"account"},"program":"spl-token"},"owner":%q}}`, mint, owner, TokenProgramIDBase58)), nil
		case "getBalance":
			return json.RawMessage(`{"context":{"slot":1},"value":1000000000}`), nil
		case "getMinimumBalanceForRentExemption":
			return json.RawMessage(`2039280`), nil
		case "getLatestBlockhash":
			return json.RawMessage(fmt.Sprintf(`{"context":{"slot":1},"value":{"blockhash":%q,"lastValidBlockHeight":123}}`, blockhash)), nil
		default:
			return nil, fmt.Errorf("unexpected RPC method %q", method)
		}
	}}

	prepared, err := client.PrepareSPLTokenTransfers(hdwallet.SolanaBase58Encode(privateKey), mint, []SolanaTransfer{
		{ToAddress: destinationOwners[0], Amount: 98_000_000},
		{ToAddress: destinationOwners[1], Amount: 2_000_000},
	})
	if err != nil {
		t.Fatalf("PrepareSPLTokenTransfers() error = %v", err)
	}
	instructions := parseTestTransferInstructions(t, prepared.RawTransaction, 11, 7)
	if len(instructions) != 3 {
		t.Fatalf("instruction count = %d, want transfer/create/transfer", len(instructions))
	}
	if instructions[0].programIDIndex != 8 || instructions[0].data[0] != 3 {
		t.Fatalf("first instruction = %#v, want existing ATA transfer", instructions[0])
	}
	if instructions[1].programIDIndex != 10 || len(instructions[1].data) != 0 {
		t.Fatalf("second instruction = %#v, want ATA creation", instructions[1])
	}
	if instructions[2].programIDIndex != 8 || instructions[2].data[0] != 3 || binary.LittleEndian.Uint64(instructions[2].data[1:]) != 2_000_000 {
		t.Fatalf("third instruction = %#v, want 2-token split transfer", instructions[2])
	}
}

func TestIsSolanaRPCErrorThroughWrapping(t *testing.T) {
	err := fmt.Errorf("send transaction error: %w", &SolanaRPCError{
		Code:    -32002,
		Message: "Transaction simulation failed",
	})
	if !IsSolanaRPCError(err) {
		t.Fatal("IsSolanaRPCError() = false, want true")
	}
	if IsSolanaRPCError(fmt.Errorf("rpc request error: timeout")) {
		t.Fatal("transport error must not be classified as an RPC rejection")
	}
}

func TestNewPreparedSolanaTransactionRejectsInvalidData(t *testing.T) {
	if _, err := newPreparedSolanaTransaction([]byte{0}, 0); err == nil {
		t.Fatal("newPreparedSolanaTransaction() error = nil, want error")
	}
}
