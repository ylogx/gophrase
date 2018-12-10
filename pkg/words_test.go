package pkg

import "testing"

func Test_ensureWordListExists_DataExists_DoNoting(t *testing.T) {
	inputVocab := []string{"hgfe", "dcba"}
	vocab := &english{wordsData: inputVocab}

	err := vocab.ensureWordListExists()
	words := vocab.words()

	if err != nil {
		t.Fatalf("Expected error to be nil but got %q", err)
	}
	if len(inputVocab) != len(words) {
		t.Fatalf("Expected %d vocab word but got %d", len(inputVocab), len(words))
	}
	for i:=0; i<len(words); i++ {
		word, inputVocabWord := words[i], inputVocab[i]
		if word != inputVocabWord {
			t.Fatalf("Expected vocab word %q but got %q", inputVocab, words)
			return
		}
	}
}

func Test_ensureWordListExists_DataExists_GiveCorrectSize(t *testing.T) {
	inputVocab := []string{"hgfe", "dcba"}
	vocab := &english{wordsData: inputVocab}

	err := vocab.ensureWordListExists()
	size := vocab.size()

	if err != nil {
		t.Fatalf("Expected error to be nil but got %q", err)
	}
	if size != len(inputVocab) {
		t.Fatalf("Expected %d vocab word but got %d", len(inputVocab), len(words))
	}
}
