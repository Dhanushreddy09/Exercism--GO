package lasagna

import "fmt"

func Lasagna() {
	noOfLayers := 5
	actualMinutesInOven := 10

	fmt.Printf("Remaining Oven time : %v\n", RemainingOvenTime(actualMinutesInOven))
	fmt.Printf("Total Preparation time : %v\n", PreparationTime(noOfLayers))
	fmt.Printf("Elapsed time is : %v\n", ElapsedTime(noOfLayers, actualMinutesInOven))
}
