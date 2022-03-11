package romannumerals

import "fmt"

func Romannumerals() {
	fmt.Println("Enter a number :")
	var input int

	fmt.Scanf("%d", &input)

	roman, err := ToRomanNumeral(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(roman)
	}
}
