package leap

import "fmt"

func Leap() {
	x := 2020
	fmt.Printf("%v is a leap year: %v\n", x, IsLeapYear(x))
}
