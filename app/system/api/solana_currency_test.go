package api

import "testing"

func TestValidateSolanaWithdrawSplit(t *testing.T) {
	validAddress := "11111111111111111111111111111111"
	tests := []struct {
		name     string
		enabled  int8
		address  string
		amount   float64
		bps      int
		decimals int
		wantErr  bool
	}{
		{name: "disabled", enabled: 0},
		{name: "fixed amount", enabled: 1, address: validAddress, amount: 2, decimals: 6},
		{name: "rate", enabled: 1, address: validAddress, bps: 200, decimals: 6},
		{name: "invalid switch", enabled: 2, wantErr: true},
		{name: "invalid address", enabled: 1, address: "invalid", amount: 2, decimals: 6, wantErr: true},
		{name: "no split mode", enabled: 1, address: validAddress, decimals: 6, wantErr: true},
		{name: "negative fixed amount", enabled: 1, address: validAddress, amount: -2, decimals: 6, wantErr: true},
		{name: "both split modes", enabled: 1, address: validAddress, amount: 2, bps: 200, decimals: 6, wantErr: true},
		{name: "rate is 100 percent", enabled: 1, address: validAddress, bps: 10000, decimals: 6, wantErr: true},
		{name: "exceeds precision", enabled: 1, address: validAddress, amount: 0.0000001, decimals: 6, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateSolanaWithdrawSplit(tt.enabled, tt.address, tt.amount, tt.bps, tt.decimals)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validateSolanaWithdrawSplit() error = %v, wantErr %t", err, tt.wantErr)
			}
		})
	}
}
