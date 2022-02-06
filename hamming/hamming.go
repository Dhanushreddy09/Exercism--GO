package hamming

import "errors"

func Distance(a, b string) (int, error) {
	runeOfA := []rune(a)
	runeOfB := []rune(b)

	if len(runeOfA) != len(runeOfB) {
		return 0, errors.New("length of strings should be same")
	}

	distance := 0
	for i, value := range runeOfA {
		if value != runeOfB[i] {
			distance++
		}
	}
	return distance, nil
}
