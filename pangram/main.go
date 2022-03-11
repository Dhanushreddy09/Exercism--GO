package pangram

import "fmt"

func Pangram() {
	fmt.Println("Enter a sentence :")
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println(IsPangram(input))
}
