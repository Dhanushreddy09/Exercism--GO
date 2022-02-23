package grains

import "fmt"

func Grains() {
	val, _ := Square(1)
	fmt.Printf("Grains present on a particular square : %d\n", val)
	fmt.Printf("Total grains on the board : %d\n", Total())
}
