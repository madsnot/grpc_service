package validators

import "unicode"

func ValidNum(str string) bool {
	for _, symb := range str {
		if !unicode.IsDigit(symb) {
			return false
		}
	}
	return true
}
