package main

import (
	"fmt"

	"github.com/jeannie/cipher-go/ciphers"
)

func main() {
	// fmt.Println("Enter your plain text:")
	// reader := bufio.NewReader(os.Stdin)
	// plainTxt, _ := reader.ReadString('\n')

	plainTxt := "This is plain text zZ123"

	cipherTxt := ciphers.GetRotThirteen(plainTxt)

	fmt.Printf("Cipher text: %s", cipherTxt)
}
