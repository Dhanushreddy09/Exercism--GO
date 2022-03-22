package word_count

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	ans := make(Frequency)
	phrase = strings.ToLower(phrase)
	words := strings.FieldsFunc(phrase, Split)
	for _, val := range words {
		val = strings.TrimLeft(strings.TrimRight(val, "'"), "'")
		_, wordExists := ans[val]
		if !wordExists {
			ans[val] = 1
		} else {
			ans[val]++
		}
	}
	return ans
}
func Split(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
}
