package isogram

import "fmt"

func Isogram() {
	var word string
	fmt.Println("Enter a word")
	fmt.Scan(&word)

	isogram := IsIsogram(word)
	fmt.Println(isogram)
}
