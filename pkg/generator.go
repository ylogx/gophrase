package pkg

import (
	"fmt"
	"math/rand"
)

func GeneratePassphrase(n int) string {
	words := ReadWords()
	wLen := len(words)
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
