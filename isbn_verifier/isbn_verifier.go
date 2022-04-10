package isbn_verifier

import (
	"unicode"
)

func IsValidISBN(isbn string) bool {
	sum := 0
	counter := 1
	digitCount := 0

	for i, val := range isbn {
		if unicode.IsDigit(val) || val == 88 || val == 120 {
			if val == 88 || val == 120 {
				if i != len(isbn)-1 {
					return false
				}
				sum += counter * 10
			} else {
				sum += counter * int(val-'0')
			}
			counter++
			digitCount++
		}
	}
	return sum%11 == 0 && digitCount == 10
}
