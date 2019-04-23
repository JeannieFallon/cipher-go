package ciphers

import (
	"strings"

	"github.com/jeannie/cipher-go/util"
)

// GetVigenere applies Vigenere cipher to given string
func GetVigenere(plainTxt string, keyWord string) string {

	var b strings.Builder
	var ascVal int
	var cipherLet string

	txtLen := len(plainTxt)
	keyLen := len(keyWord)

	// get ascVal of each letter in keyword
	shiftVals := util.GetShiftVals(keyWord, keyLen)

	// loop over each letter of plaintext
	for i := 0; i < txtLen; i++ {

		// loop over each letter of keyword
		for j := 0; j < keyLen; j++ {

			if i < len(plainTxt) {
				ascVal = int(plainTxt[i])
				keyVal := shiftVals[j]
				cipherLet = util.GetCipherLetter(ascVal, keyVal)
				b.WriteString(cipherLet)

				i++

			} else {
				break
			}
		}
	}

	return b.String()
}
