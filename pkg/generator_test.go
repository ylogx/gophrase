package pkg

import (
	"testing"
)

func TestEmptyPassphrase(t *testing.T) {
	passphrase := GeneratePassphrase()

	if passphrase != "" {
		t.Fatalf("Expected passphrase to be empty but was %s", passphrase)
	}
}
