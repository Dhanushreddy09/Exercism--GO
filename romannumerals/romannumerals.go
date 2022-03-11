package romannumerals

import "errors"

type Roman struct {
	roman   string
	integer int
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3000 {
		return "", errors.New("entered input is beyond the roman numbers range")
	}
	romans := []Roman{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"XXX", 30},
		{"XX", 20},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	var ans string

	for _, val := range romans {
		x := val.integer
		for input >= x {
			input -= x
			ans += val.roman
		}
	}
	return ans, nil
}
