package util

import "strings"

// GetCipherLetter returns
func GetCipherLetter(ascVal int, shiftVal int) string {

	// only shift upper- and lowercase letters
	if ascVal >= EngUpperFloor && ascVal <= EngUpperCeiling {
		ascVal = GetCipherASCII(EngUpperFloor, ascVal, shiftVal)
	} else if ascVal >= EngLowerFloor && ascVal <= EngLowerCeiling {
		ascVal = GetCipherASCII(EngLowerFloor, ascVal, shiftVal)
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
	return (alphaIdx + shiftVal) % EngAlphaLen
}

// GetShiftVals returns shift values from letters of keyword.
func GetShiftVals(keyWord string, keyLen int) []int {

	var ascVal int
	shiftVals := make([]int, keyLen)
	keyWordLower := strings.ToLower(keyWord)

	for i := 0; i < keyLen; i++ {
		ascVal = int(keyWordLower[i])
		shiftVals[i] = GetShiftVal(ascVal)
	}

	return shiftVals
}

// GetShiftVal returns shift value for a letter based upon alphabet index plus one.
func GetShiftVal(ascVal int) int {
	// only pass lowercase letters to normalize ASCII to index translation
	return ((ascVal - EngLowerFloor) % 26) + 1
}
