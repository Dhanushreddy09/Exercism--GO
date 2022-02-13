package card_tricks

import "fmt"

func Card_tricks() {
	slice := []int{5, 6, 8, 10, 12}
	fmt.Println(GetItem(slice, 2))
	fmt.Println(SetItem(slice, 8, 7))
}
