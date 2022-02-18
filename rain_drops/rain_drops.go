package rain_drops

import "fmt"

func Convert(number int) string {
	var raindrops string
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		return fmt.Sprint(number)
	}
	if number%3 == 0 {
		raindrops += "Pling"
	}
	if number%5 == 0 {
		raindrops += "Plang"
	}
	if number%7 == 0 {
		raindrops += "Plong"
	}

	return raindrops
}
