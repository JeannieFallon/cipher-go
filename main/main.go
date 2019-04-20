package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jeannie/cipher-go/ciphers"
)

func main() {
	fmt.Println("Enter your plain text:")
	reader := bufio.NewReader(os.Stdin)
	plainTxt, _ := reader.ReadString('\n')
	cipherTxt := ciphers.GetRotThirteen(plainTxt)
	fmt.Printf("Rot-13 cipher text:\n%s", cipherTxt)
}
