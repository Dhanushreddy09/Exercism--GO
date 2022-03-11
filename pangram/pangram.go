package pangram

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	alphabets := map[rune]bool{}
	for _, val := range input {
		if !unicode.IsLetter(val) {
			continue
		}
		alphabets[val] = true
	}
	fmt.Println(alphabets)
	return len(alphabets) == 26
}
