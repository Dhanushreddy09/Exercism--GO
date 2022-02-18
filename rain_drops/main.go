package rain_drops

import "fmt"

func Rain_drops() {
	var number int
	fmt.Println("Enter a number")
	fmt.Scan(&number)

	raindrops := Convert(number)
	fmt.Println(raindrops)
}
