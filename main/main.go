package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jeannie/cipher-go/ciphers"
	"github.com/jeannie/cipher-go/util"
)

func main() {
	fmt.Println("Enter your plain text:")
	reader := bufio.NewReader(os.Stdin)
	plainTxt, _ := reader.ReadString('\n')
	fmt.Printf("Cipher text: %s", plainTxt)

	ciphers.GetRotThirteen()
	ciphers.GetCaesar()
	ciphers.GetVigenere()

	util.GetAlphaShift()
}
