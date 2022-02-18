package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	letterSlice := []rune{}
	wordRune := []rune(word)

	for i := 0; i < len(wordRune); i++ {
		for _, value := range letterSlice {
			if wordRune[i] == value && unicode.IsLetter(value) {
				return false
			}
		}
		letterSlice = append(letterSlice, wordRune[i])
	}
	return true
}
