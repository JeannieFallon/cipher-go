package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jeannie/cipher-go/ciphers"
)

// TODO validation

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your plain text:")
	plainTxt, _ := reader.ReadString('\n')

	fmt.Println("Enter a key value for the Caesar cipher:")
	keyStr, _ := reader.ReadString('\n')
	keyVal, _ := strconv.Atoi(keyStr)

	rot13Cipher := ciphers.GetRotThirteen(plainTxt)
	caesarCipher := ciphers.GetCaesar(plainTxt, keyVal)

	fmt.Printf("Rot-13 cipher text:\n%s", rot13Cipher)
	fmt.Printf("Caesar cipher text is:\n%s", caesarCipher)
}
