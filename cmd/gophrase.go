package main

import (
	"../pkg"
	"fmt"
	"flag"
	"math/rand"
	"time"
)

const (
	DefaultLength = 4
)

func main() {
	numWords := flag.Int("n", DefaultLength, "Number of words in passphrase")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	passphrase := pkg.GeneratePassphrase(*numWords)
	fmt.Printf("Passphrase with %d words and length of %d chars:\n%s\n", *numWords, len(passphrase), passphrase)
}
