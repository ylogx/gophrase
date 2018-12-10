package pkg

import (
	"fmt"
	"math/rand"
)

func GeneratePassphrase(n int, vocab Vocabulary) string {
	words := vocab.words()
	wLen := vocab.size()
	if wLen <= 0 {
		panic("No word list found!")
	}
	output := ""
	for i := 0; i < n; i++ {
		index := rand.Intn(wLen)
		//fmt.Printf("Index: %d %q\n", index, words[i])
		output += string(words[index])
		if i != n-1 {
			output += " "
		}
		//fmt.Printf("Output: %s\n", output)
	}
	return fmt.Sprintf("%s", output)
}
