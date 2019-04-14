package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your plain text:")
	reader := bufio.NewReader(os.Stdin)
	plainTxt, _ := reader.ReadString('\n')
	fmt.Printf("Cipher text: %s", plainTxt)
}
