package main

import (
	"exercism/cars_assemble"
	"exercism/hamming"
	"exercism/lasagna"
	"exercism/twofer"
	"fmt"
)

func main() {
	//Call respective functions here to find out the output of respective problems

	//Note: Function names are same as packages names or folder names in which the function is present

	/*But don't forget that first letter of the function name should be capitalized
	since we cannot use same name as packages */

	//Below is an example for your reference

	Cars_assemble()
}

func Twofer() {
	var name string

	fmt.Println("Enter name : ")
	fmt.Scanf("%s", &name)

	fmt.Println(twofer.ShareWith(name))
}

func Lasagna() {
	noOfLayers := 5
	actualMinutesInOven := 10

	fmt.Printf("Remaining Oven time : %v\n", lasagna.RemainingOvenTime(actualMinutesInOven))
	fmt.Printf("Total Preparation time : %v\n", lasagna.PreparationTime(noOfLayers))
	fmt.Printf("Elapsed time is : %v\n", lasagna.ElapsedTime(noOfLayers, actualMinutesInOven))
}

func Hamming() {
	var a, b string

	fmt.Println("Enter first string : ")
	fmt.Scan(&a)

	fmt.Println("Enter second string : ")
	fmt.Scan(&b)

	distance, err := hamming.Distance(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(distance)
	}
}

func Cars_assemble() {
	var productionRate int
	var successRate float64

	fmt.Println("Enter production rate : ")
	fmt.Scan(&productionRate)

	fmt.Println("Enter success rate : ")
	fmt.Scan(&successRate)

	fmt.Println("Working cars per hour : ", cars_assemble.CalculateCost(148))
}
