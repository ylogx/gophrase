package pkg

import (
	"math/rand"
	"testing"
)

var vocab = &english{wordsData: []string{"abcd", "efgh", "ijkl", "mnop"}}

func TestEmptyPassphrase(t *testing.T) {
	rand.Seed(26)
	passphrase := GeneratePassphrase(0, vocab)

	if passphrase != "" {
		t.Fatalf("Expected passphrase to be empty but was '%s'", passphrase)
	}
}

func TestWordsIsFromVocabPassphrase(t *testing.T) {
	rand.Seed(26)
	expected := "efgh"

	passphrase := GeneratePassphrase(1, vocab)

	if passphrase != expected {
		t.Fatalf("Expected passphrase to be '%s' but was '%s'", expected, passphrase)
	}
}

func TestWordsAreInLimitPassphrase(t *testing.T) {
	rand.Seed(26)
	passphrase := GeneratePassphrase(1, vocab)

	if WordMinLength > len(passphrase) || len(passphrase) > WordMaxLength {
		t.Fatalf("Expected passphrase to be %d-%d char long but was '%s'", WordMinLength, WordMaxLength, passphrase)
	}
}

func Test6WordsIsFromVocabPassphrase(t *testing.T) {
	rand.Seed(26)
	expected := "efgh mnop mnop efgh efgh mnop"

	passphrase := GeneratePassphrase(6, vocab)

	if passphrase != expected {
		t.Fatalf("Expected passphrase to be '%s' but was '%s'", expected, passphrase)
	}
}

func TestMultiPhraseWordsAreInLimitPassphrase(t *testing.T) {
	rand.Seed(26)
	passphrase := GeneratePassphrase(6, vocab)

	if 6*WordMinLength > len(passphrase) || len(passphrase) > 6*WordMaxLength {
		t.Fatalf("Expected passphrase to be %d-%d char long but was '%s'", 6*WordMinLength, 6*WordMaxLength, passphrase)
	}
}
