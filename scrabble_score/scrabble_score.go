package scrabble_score

import "strings"

func Score(word string) int {
	word = strings.ToLower(word)
	scrabbleScore := 0

	for i := 0; i < len(word); i++ {
		switch word[i] {
		case 'k':
			scrabbleScore += 5
		case 'd', 'g':
			scrabbleScore += 2
		case 'b', 'c', 'm', 'p':
			scrabbleScore += 3
		case 'f', 'h', 'v', 'w', 'y':
			scrabbleScore += 4
		case 'j', 'x':
			scrabbleScore += 8
		case 'q', 'z':
			scrabbleScore += 10
		default:
			scrabbleScore += 1
		}
	}
	return scrabbleScore
}
