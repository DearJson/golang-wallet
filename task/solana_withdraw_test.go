package task

import (
	"testing"

	"gfast/app/system/model"
)

func TestBuildSolanaWithdrawTransfersDisabled(t *testing.T) {
	withdrawal := &model.Withdraw{Address: "user-address", Amount: 100}
	currency := &model.Currency{Decimals: 6}

	transfers, splitAddress, splitAmount, err := buildSolanaWithdrawTransfers(withdrawal, currency)
	if err != nil {
		t.Fatalf("buildSolanaWithdrawTransfers() error = %v", err)
	}
	if len(transfers) != 1 || transfers[0].ToAddress != withdrawal.Address || transfers[0].Amount != 100_000_000 {
		t.Fatalf("transfers = %#v, want one 100-token user transfer", transfers)
	}
	if splitAddress != "" || splitAmount != "0" {
		t.Fatalf("split snapshot = (%q, %q), want empty address and zero amount", splitAddress, splitAmount)
	}
}

func TestBuildSolanaWithdrawTransfersFixedSplit(t *testing.T) {
	withdrawal := &model.Withdraw{Address: "user-address", Amount: 100}
	currency := &model.Currency{
		Name:                 "USDT",
		Decimals:             6,
		WithdrawSplitEnabled: 1,
		WithdrawSplitAddress: "fixed-address",
		WithdrawSplitAmount:  "2",
	}

	transfers, splitAddress, splitAmount, err := buildSolanaWithdrawTransfers(withdrawal, currency)
	if err != nil {
		t.Fatalf("buildSolanaWithdrawTransfers() error = %v", err)
	}
	if len(transfers) != 2 {
		t.Fatalf("len(transfers) = %d, want 2", len(transfers))
	}
	if transfers[0].ToAddress != withdrawal.Address || transfers[0].Amount != 98_000_000 {
		t.Fatalf("user transfer = %#v, want 98 tokens", transfers[0])
	}
	if transfers[1].ToAddress != currency.WithdrawSplitAddress || transfers[1].Amount != 2_000_000 {
		t.Fatalf("split transfer = %#v, want 2 tokens", transfers[1])
	}
	if splitAddress != currency.WithdrawSplitAddress || splitAmount != "2" {
		t.Fatalf("split snapshot = (%q, %q), want (%q, 2)", splitAddress, splitAmount, currency.WithdrawSplitAddress)
	}
}

func TestBuildSolanaWithdrawTransfersRejectsInvalidSplit(t *testing.T) {
	tests := []struct {
		name       string
		withdrawal *model.Withdraw
		currency   *model.Currency
	}{
		{
			name:       "split equals total",
			withdrawal: &model.Withdraw{Address: "user-address", Amount: 2},
			currency:   &model.Currency{Name: "USDT", Decimals: 6, WithdrawSplitEnabled: 1, WithdrawSplitAddress: "fixed-address", WithdrawSplitAmount: "2"},
		},
		{
			name:       "same recipient",
			withdrawal: &model.Withdraw{Address: "same-address", Amount: 100},
			currency:   &model.Currency{Name: "USDT", Decimals: 6, WithdrawSplitEnabled: 1, WithdrawSplitAddress: "same-address", WithdrawSplitAmount: "2"},
		},
		{
			name:       "less than base unit",
			withdrawal: &model.Withdraw{Address: "user-address", Amount: 100},
			currency:   &model.Currency{Name: "USDT", Decimals: 6, WithdrawSplitEnabled: 1, WithdrawSplitAddress: "fixed-address", WithdrawSplitAmount: "0.0000001"},
		},
		{
			name:       "invalid switch",
			withdrawal: &model.Withdraw{Address: "user-address", Amount: 100},
			currency:   &model.Currency{Name: "USDT", Decimals: 6, WithdrawSplitEnabled: 2, WithdrawSplitAddress: "fixed-address", WithdrawSplitAmount: "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, _, _, err := buildSolanaWithdrawTransfers(tt.withdrawal, tt.currency); err == nil {
				t.Fatal("buildSolanaWithdrawTransfers() error = nil, want error")
			}
		})
	}
}
