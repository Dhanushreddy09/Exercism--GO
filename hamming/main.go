package hamming

import "fmt"

func Hamming() {
	var a, b string

	fmt.Println("Enter first string : ")
	fmt.Scan(&a)

	fmt.Println("Enter second string : ")
	fmt.Scan(&b)

	distance, err := Distance(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(distance)
	}
}
