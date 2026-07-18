package main

import (
	"bufio"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"gfast/hdwallet"
	"gfast/library"
)

func decryptSolanaPrivateKey(reader io.Reader, writer io.Writer) error {
	encrypted, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("read encrypted private key: %w", err)
	}

	decrypted, err := library.DecryptByAes(strings.TrimSpace(string(encrypted)))
	if err != nil {
		return fmt.Errorf("decrypt private key: %w", err)
	}
	privateKey := string(decrypted)
	if err = validateSolanaPrivateKey(privateKey); err != nil {
		return fmt.Errorf("invalid Solana private key: %w", err)
	}

	_, err = fmt.Fprintln(writer, privateKey)
	return err
}

func validateSolanaPrivateKey(privateKey string) error {
	if decoded, err := hex.DecodeString(privateKey); err == nil && len(decoded) == ed25519.PrivateKeySize {
		return nil
	}

	decoded, err := hdwallet.SolanaBase58Decode(privateKey)
	if err != nil {
		return fmt.Errorf("expected a 64-byte hex or Base58 private key: %w", err)
	}
	if len(decoded) != ed25519.PrivateKeySize {
		return fmt.Errorf("invalid private key length: %d, expected: %d", len(decoded), ed25519.PrivateKeySize)
	}
	return nil
}

func main() {
	reader := io.Reader(os.Stdin)
	if stat, err := os.Stdin.Stat(); err == nil && stat.Mode()&os.ModeCharDevice != 0 {
		fmt.Fprintln(os.Stderr, "Paste the encrypted Solana private key, then press Enter:")
		line, readErr := bufio.NewReader(os.Stdin).ReadString('\n')
		if readErr != nil && readErr != io.EOF {
			fmt.Fprintf(os.Stderr, "Error: read encrypted private key: %v\n", readErr)
			os.Exit(1)
		}
		reader = strings.NewReader(line)
	}

	if err := decryptSolanaPrivateKey(reader, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
