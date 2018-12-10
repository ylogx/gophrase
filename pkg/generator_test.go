package pkg

import (
	"testing"
)

func TestEmptyPassphrase(t *testing.T) {
	passphrase := GeneratePassphrase(0)

	if passphrase != "" {
		t.Fatalf("Expected passphrase to be empty but was '%s'", passphrase)
	}
}

func TestWordsAreInLimitPassphrase(t *testing.T) {
	passphrase := GeneratePassphrase(1)

	if WordMinLength > len(passphrase) || len(passphrase) > WordMaxLength {
		t.Fatalf("Expected passphrase to be %d-%d char long but was '%s'", WordMinLength, WordMaxLength, passphrase)
	}
}

func TestMultiPhraseWordsAreInLimitPassphrase(t *testing.T) {
	passphrase := GeneratePassphrase(6)

	if 6*WordMinLength > len(passphrase) || len(passphrase) > 6*WordMaxLength {
		t.Fatalf("Expected passphrase to be %d-%d char long but was '%s'", 6*WordMinLength, 6*WordMaxLength, passphrase)
	}
}
