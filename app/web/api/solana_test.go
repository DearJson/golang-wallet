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

func TestParseSolanaContractDepositInstructionDepositUsdtReserve(t *testing.T) {
	orderID := "ORD-RESERVE-001"
	data := make([]byte, 13+len(orderID))
	data[0] = solanaDepositUsdtReserveInstructionIndex
	binary.LittleEndian.PutUint64(data[1:9], 1000000)
	binary.LittleEndian.PutUint32(data[9:13], uint32(len(orderID)))
	copy(data[13:], orderID)

	got, err := parseSolanaContractDepositInstruction(data)
	if err != nil {
		t.Fatalf("parseSolanaContractDepositInstruction() error = %v", err)
	}

	if got.Index != solanaDepositUsdtReserveInstructionIndex {
		t.Fatalf("Index = %d, want %d", got.Index, solanaDepositUsdtReserveInstructionIndex)
	}
	if got.Amount != 1000000 {
		t.Fatalf("Amount = %d, want 1000000", got.Amount)
	}
	if got.OrderId != orderID {
		t.Fatalf("OrderId = %q, want %q", got.OrderId, orderID)
	}
}

func TestParseSolanaContractDepositInstructionDepositUsdtSplit(t *testing.T) {
	orderID := "ORD-SPLIT-001"
	recipients := []string{
		"9inG5NgrGUpxo6htR2aKSZXGeYybtg9QaC6iwRNn77D5",
		"3QqJ9HPYVGb16WzjZ8xWUxEvNPd46Xdm8tYcAE3MLAE6",
	}
	amounts := []uint64{270000000, 240000000}
	data := buildDepositUsdtSplitInstruction(t, 510000000, recipients, amounts, orderID)

	got, err := parseSolanaContractDepositInstruction(data)
	if err != nil {
		t.Fatalf("parseSolanaContractDepositInstruction() error = %v", err)
	}

	if got.Index != solanaDepositUsdtSplitInstructionIndex {
		t.Fatalf("Index = %d, want %d", got.Index, solanaDepositUsdtSplitInstructionIndex)
	}
	if got.Amount != 510000000 {
		t.Fatalf("Amount = %d, want 510000000", got.Amount)
	}
	if got.OrderId != orderID {
		t.Fatalf("OrderId = %q, want %q", got.OrderId, orderID)
	}
	if len(got.Recipients) != len(recipients) {
		t.Fatalf("Recipients length = %d, want %d", len(got.Recipients), len(recipients))
	}
	for i := range recipients {
		if got.Recipients[i] != recipients[i] {
			t.Fatalf("Recipients[%d] = %q, want %q", i, got.Recipients[i], recipients[i])
		}
		if got.Amounts[i] != amounts[i] {
			t.Fatalf("Amounts[%d] = %d, want %d", i, got.Amounts[i], amounts[i])
		}
	}
}

func TestParseSolanaContractDepositInstructionRejectsSplitSumMismatch(t *testing.T) {
	recipients := []string{"9inG5NgrGUpxo6htR2aKSZXGeYybtg9QaC6iwRNn77D5"}
	amounts := []uint64{999999}
	data := buildDepositUsdtSplitInstruction(t, 1000000, recipients, amounts, "ORD-SPLIT-BAD")

	if _, err := parseSolanaContractDepositInstruction(data); err == nil {
		t.Fatal("parseSolanaContractDepositInstruction() error = nil, want split sum mismatch error")
	}
}

func TestParseSolanaContractDepositInstructionRejectsLongReserveOrder(t *testing.T) {
	orderID := "123456789012345678901234567890123"
	data := make([]byte, 13+len(orderID))
	data[0] = solanaDepositUsdtReserveInstructionIndex
	binary.LittleEndian.PutUint64(data[1:9], 1)
	binary.LittleEndian.PutUint32(data[9:13], uint32(len(orderID)))
	copy(data[13:], orderID)

	if _, err := parseSolanaContractDepositInstruction(data); err == nil {
		t.Fatal("parseSolanaContractDepositInstruction() error = nil, want long reserve order error")
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

func TestExtractDepositUsdtSplitTransferAmount(t *testing.T) {
	depositIx := &solanaContractDepositInstruction{
		Index:      solanaDepositUsdtSplitInstructionIndex,
		Amount:     536000000,
		Recipients: []string{"user-change", "recipient-1", "recipient-2", "recipient-3"},
		Amounts:    []uint64{270000000, 240000000, 16000000, 10000000},
		OrderId:    "ORD-SPLIT-001",
	}
	tx := &rpc.HeliusEnhancedTransaction{
		TokenTransfers: []rpc.HeliusTokenTransfer{
			{
				FromTokenAccount: "user-usdt-ata",
				FromUserAccount:  "user",
				Mint:             "usdt-mint",
				ToTokenAccount:   "user-change-usdt-ata",
				ToUserAccount:    "user",
				TokenAmount:      270,
			},
			{
				FromTokenAccount: "user-usdt-ata",
				FromUserAccount:  "user",
				Mint:             "usdt-mint",
				ToTokenAccount:   "recipient-1-usdt-ata",
				ToUserAccount:    "recipient-1",
				TokenAmount:      240,
			},
			{
				FromTokenAccount: "user-usdt-ata",
				FromUserAccount:  "user",
				Mint:             "usdt-mint",
				ToTokenAccount:   "recipient-2-usdt-ata",
				ToUserAccount:    "recipient-2",
				TokenAmount:      16,
			},
			{
				FromTokenAccount: "user-usdt-ata",
				FromUserAccount:  "user",
				Mint:             "usdt-mint",
				ToTokenAccount:   "recipient-3-usdt-ata",
				ToUserAccount:    "recipient-3",
				TokenAmount:      10,
			},
		},
		AccountData: []rpc.HeliusAccountData{
			newTokenBalanceChangeAccountData("user-change-usdt-ata", "usdt-mint", 6, "0"),
			newTokenBalanceChangeAccountData("recipient-1-usdt-ata", "usdt-mint", 6, "240000000"),
			newTokenBalanceChangeAccountData("recipient-2-usdt-ata", "usdt-mint", 6, "16000000"),
			newTokenBalanceChangeAccountData("recipient-3-usdt-ata", "usdt-mint", 6, "10000000"),
		},
	}

	mint, amount, err := extractDepositUsdtSplitTransferAmount(tx, "user", "user-usdt-ata", []string{"user-change-usdt-ata", "recipient-1-usdt-ata", "recipient-2-usdt-ata", "recipient-3-usdt-ata"}, depositIx)
	if err != nil {
		t.Fatalf("extractDepositUsdtSplitTransferAmount() error = %v", err)
	}
	if mint != "usdt-mint" {
		t.Fatalf("mint = %q, want usdt-mint", mint)
	}
	if amount.String() != "266" {
		t.Fatalf("amount = %s, want 266", amount.String())
	}

	if _, _, err := extractDepositUsdtSplitTransferAmount(tx, "user", "user-usdt-ata", []string{"user-change-usdt-ata", "recipient-1-usdt-ata", "recipient-2-usdt-ata", "missing-recipient-usdt-ata"}, depositIx); err == nil {
		t.Fatal("extractDepositUsdtSplitTransferAmount() error = nil for missing recipient, want error")
	}
}

func buildDepositUsdtSplitInstruction(t *testing.T, amount uint64, recipients []string, amounts []uint64, orderID string) []byte {
	t.Helper()

	data := make([]byte, 0, 1+8+4+len(recipients)*32+4+len(amounts)*8+4+len(orderID))
	data = append(data, solanaDepositUsdtSplitInstructionIndex)
	data = binary.LittleEndian.AppendUint64(data, amount)
	data = binary.LittleEndian.AppendUint32(data, uint32(len(recipients)))
	for _, recipient := range recipients {
		pubkey, err := decodeSolanaInstructionData(recipient)
		if err != nil {
			t.Fatalf("decodeSolanaInstructionData(%q) error = %v", recipient, err)
		}
		if len(pubkey) != 32 {
			t.Fatalf("recipient pubkey length = %d, want 32", len(pubkey))
		}
		data = append(data, pubkey...)
	}
	data = binary.LittleEndian.AppendUint32(data, uint32(len(amounts)))
	for _, amount := range amounts {
		data = binary.LittleEndian.AppendUint64(data, amount)
	}
	data = binary.LittleEndian.AppendUint32(data, uint32(len(orderID)))
	data = append(data, orderID...)
	return data
}

func newTokenBalanceChangeAccountData(tokenAccount string, mint string, decimals uint8, tokenAmount string) rpc.HeliusAccountData {
	return rpc.HeliusAccountData{
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
				Mint:         mint,
				TokenAccount: tokenAccount,
				RawTokenAmount: struct {
					Decimals    uint8  `json:"decimals"`
					TokenAmount string `json:"tokenAmount"`
				}{
					Decimals:    decimals,
					TokenAmount: tokenAmount,
				},
			},
		},
	}
}
