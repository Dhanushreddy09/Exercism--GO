package cars_assemble

import "fmt"

func Cars_assemble() {
	var productionRate int
	var successRate float64

	fmt.Println("Enter production rate : ")
	fmt.Scan(&productionRate)

	fmt.Println("Enter success rate : ")
	fmt.Scan(&successRate)

	fmt.Println("Working cars per hour : ", CalculateCost(148))
}
