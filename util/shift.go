package util

//TODO hardcoded values into constants

// GetCipherLetter returns
func GetCipherLetter(ascVal int, shiftVal int) string {

	// only shift upper- and lowercase letters
	if ascVal >= 65 && ascVal <= 90 {
		ascVal = GetCipherASCII(65, ascVal, shiftVal)
	} else if ascVal >= 97 && ascVal <= 122 {
		ascVal = GetCipherASCII(97, ascVal, shiftVal)
	}

	// convert ASCII val back to letter
	return string(ascVal)
}

// GetCipherASCII returns new ASCII value of shifted letter.
func GetCipherASCII(alphaMapVal int, ascVal int, shiftVal int) int {
	// map ASCII value onto index of lower- or uppercase letter in alphabet
	alphaIdx := ascVal - alphaMapVal
	// wrap shift rotation around alphabet indices
	alphaIdx = GetAlphaIndex(alphaIdx, shiftVal)
	// map alphabet index back onto original ASCII range
	return alphaIdx + alphaMapVal
}

// GetAlphaIndex applies shift value and returns cipher letter index.
func GetAlphaIndex(alphaIdx int, shiftVal int) int {
	return (alphaIdx + shiftVal) % 26
}
