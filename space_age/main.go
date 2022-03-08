package space_age

import "fmt"

func Space_age() {
	var age float64
	var planet Planet
	fmt.Print("Enter your age in seconds: ")
	fmt.Scan(&age)
	fmt.Print("Enter your planet(first letter should be capital): ")
	fmt.Scan(&planet)
	fmt.Println(Age(age, planet))
}
