package luhn

import (
	"strconv"
)

func Valid(id string) bool {
	if len(id) <= 1 {
		return false
	}
	idRune := []rune(id)
	counter := 0
	var revisedString string
	onlyZeroes := true
	noOfZeroes := 0
	var sum int

	for i := len(idRune) - 1; i >= 0; i-- {
		val := idRune[i]
		if val != 32 && (val < 48 || val > 57) {
			return false
		}
		if val == 32 {
			continue
		}

		if onlyZeroes {
			if val == 48 {
				noOfZeroes++
			} else {
				onlyZeroes = !onlyZeroes
			}
		}
		if counter%2 != 0 {
			revisedVal := (int(val - 48)) * 2
			if revisedVal > 9 {
				revisedVal -= 9
			}
			revisedString += strconv.Itoa(revisedVal)
		} else {
			revisedString += string(val)
		}
		counter++
	}

	if onlyZeroes && noOfZeroes <= 1 {
		return false
	}
	revisedStringRune := []rune(revisedString)
	for _, value := range revisedStringRune {
		val := int(value - 48)
		sum += val
	}
	return sum%10 == 0
}
