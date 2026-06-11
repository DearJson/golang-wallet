package api

import (
	"encoding/binary"
	"gfast/rpc"
	"testing"
)

func TestParseSolanaContractDepositInstructionUSDT(t *testing.T) {
	orderID := "ORD-USDT-001"
	data := make([]byte, 45+len(orderID))
	data[0] = solanaDepositInstructionIndex
	binary.LittleEndian.PutUint64(data[1:9], 123456)
	binary.LittleEndian.PutUint32(data[41:45], uint32(len(orderID)))
	copy(data[45:], orderID)

	got, err := parseSolanaContractDepositInstruction(data)
	if err != nil {
		t.Fatalf("parseSolanaContractDepositInstruction() error = %v", err)
	}

	if got.Index != solanaDepositInstructionIndex {
		t.Fatalf("Index = %d, want %d", got.Index, solanaDepositInstructionIndex)
	}
	if got.Amount != 123456 {
		t.Fatalf("Amount = %d, want 123456", got.Amount)
	}
	if got.OrderId != orderID {
		t.Fatalf("OrderId = %q, want %q", got.OrderId, orderID)
	}
}

func TestParseSolanaContractDepositInstructionPythia(t *testing.T) {
	orderID := "ORD-PYTHIA-001"
	data := make([]byte, 13+len(orderID))
	data[0] = solanaDepositPythiaInstructionIndex
	binary.LittleEndian.PutUint64(data[1:9], 987654)
	binary.LittleEndian.PutUint32(data[9:13], uint32(len(orderID)))
	copy(data[13:], orderID)

	got, err := parseSolanaContractDepositInstruction(data)
	if err != nil {
		t.Fatalf("parseSolanaContractDepositInstruction() error = %v", err)
	}

	if got.Index != solanaDepositPythiaInstructionIndex {
		t.Fatalf("Index = %d, want %d", got.Index, solanaDepositPythiaInstructionIndex)
	}
	if got.Amount != 987654 {
		t.Fatalf("Amount = %d, want 987654", got.Amount)
	}
	if got.OrderId != orderID {
		t.Fatalf("OrderId = %q, want %q", got.OrderId, orderID)
	}
}

func TestParseSolanaContractDepositInstructionDepositV5(t *testing.T) {
	orderID := "ORD-V5-001"
	data := make([]byte, 109+len(orderID))
	data[0] = solanaDepositV5InstructionIndex
	binary.LittleEndian.PutUint64(data[1:9], 1234567)
	binary.LittleEndian.PutUint32(data[105:109], uint32(len(orderID)))
	copy(data[109:], orderID)

	got, err := parseSolanaContractDepositInstruction(data)
	if err != nil {
		t.Fatalf("parseSolanaContractDepositInstruction() error = %v", err)
	}

	if got.Index != solanaDepositV5InstructionIndex {
		t.Fatalf("Index = %d, want %d", got.Index, solanaDepositV5InstructionIndex)
	}
	if got.Amount != 1234567 {
		t.Fatalf("Amount = %d, want 1234567", got.Amount)
	}
	if got.OrderId != orderID {
		t.Fatalf("OrderId = %q, want %q", got.OrderId, orderID)
	}
}

func TestParseSolanaContractDepositInstructionRejectsTruncatedPythiaOrder(t *testing.T) {
	data := make([]byte, 13)
	data[0] = solanaDepositPythiaInstructionIndex
	binary.LittleEndian.PutUint64(data[1:9], 1)
	binary.LittleEndian.PutUint32(data[9:13], 4)

	if _, err := parseSolanaContractDepositInstruction(data); err == nil {
		t.Fatal("parseSolanaContractDepositInstruction() error = nil, want truncated order error")
	}
}

func TestHasMatchingPythiaTransfer(t *testing.T) {
	tx := &rpc.HeliusEnhancedTransaction{
		TokenTransfers: []rpc.HeliusTokenTransfer{
			{
				FromTokenAccount: "user-pythia-ata",
				FromUserAccount:  "user",
				Mint:             solanaPythiaMint,
				ToTokenAccount:   "target-pythia-ata",
				ToUserAccount:    solanaPythiaTargetOwner,
			},
		},
		AccountData: []rpc.HeliusAccountData{
			{
				TokenBalanceChanges: []struct {
					Mint           string `json:"mint"`
					RawTokenAmount struct {
						Decimals    uint8  `json:"decimals"`
						TokenAmount string `json:"tokenAmount"`
					} `json:"rawTokenAmount"`
					TokenAccount string `json:"tokenAccount"`
					UserAccount  string `json:"userAccount"`
				}{
					{
						Mint:         solanaPythiaMint,
						TokenAccount: "target-pythia-ata",
						UserAccount:  "target",
						RawTokenAmount: struct {
							Decimals    uint8  `json:"decimals"`
							TokenAmount string `json:"tokenAmount"`
						}{
							Decimals:    6,
							TokenAmount: "1000000",
						},
					},
				},
			},
		},
	}

	if !hasMatchingPythiaTransfer(tx, "user", "user-pythia-ata", "target-pythia-ata", 1000000) {
		t.Fatal("hasMatchingPythiaTransfer() = false, want true")
	}
	if hasMatchingPythiaTransfer(tx, "user", "user-pythia-ata", "other-target", 1000000) {
		t.Fatal("hasMatchingPythiaTransfer() = true for wrong target, want false")
	}
	if hasMatchingPythiaTransfer(tx, "user", "user-pythia-ata", "target-pythia-ata", 999999) {
		t.Fatal("hasMatchingPythiaTransfer() = true for wrong amount, want false")
	}
}
