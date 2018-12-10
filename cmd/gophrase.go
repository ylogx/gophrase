package main
import (
	"../pkg"
	"fmt"
	"flag"
	"math/rand"
	"time"
)

func main() {
	numWords := flag.Int("n", 6, "Number of words in passphrase")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	passphrase := pkg.GeneratePassphrase(*numWords)
	fmt.Printf("Passphrase of length %d:\n%s\n", len(passphrase), passphrase)
}
