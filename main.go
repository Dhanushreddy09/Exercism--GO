package main

import (
	"exercism/annalyn"
	"exercism/cars_assemble"
	"exercism/hamming"
	"exercism/lasagna"
	"exercism/party_robot"
	"exercism/tech_palace"
	"exercism/twofer"
	"exercism/vehicle_purchase"
	"fmt"
)

func main() {
	//Call respective functions here to find out the output of respective problems

	//Note: Function names are same as packages names or folder names in which the function is present

	/*But don't forget that first letter of the function name should be capitalized
	since we cannot use same name as packages */

	//Below is an example for your reference

	Party_robot()
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

func Annalyn() {
	var knightIsAwake, archerIsAwake, prisonerIsAwake bool

	fmt.Println("Enter knightIsAwake : ")
	fmt.Scan(&knightIsAwake)

	fmt.Println("Enter archerIsAwake : ")
	fmt.Scan(&archerIsAwake)

	fmt.Println("Enter prisonerIsAwake : ")
	fmt.Scan(&prisonerIsAwake)

	fmt.Println("Can Fast Attack : ", annalyn.CanFastAttack(knightIsAwake))
	fmt.Println("Can Spy : ", annalyn.CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	fmt.Println("Can Signal Prisoner : ", annalyn.CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
	fmt.Println("Can Free Prisoner : ", annalyn.CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, false))
}

func Tech_palace() {
	var customer string
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

	fmt.Println("Enter customer : ")
	fmt.Scan(&customer)

	fmt.Println(tech_palace.WelcomeMessage(customer))
	fmt.Println(tech_palace.AddBorder(tech_palace.WelcomeMessage(customer), 5))
	fmt.Println(tech_palace.CleanupMessage(message))
}

func Vehicle_purchase() {
	var kind, option1, option2 string

	fmt.Println("Enter kind of vehicle : ")
	fmt.Scan(&kind)

	fmt.Println("Enter option 1 : ")
	fmt.Scan(&option1)

	fmt.Println("Enter option2 : ")
	fmt.Scan(&option2)

	fmt.Println(vehicle_purchase.NeedsLicense(kind))
	fmt.Println(vehicle_purchase.ChooseVehicle(option1, option2))
	fmt.Println(vehicle_purchase.CalculateResellPrice(10000.0, 5.0))
}

func Party_robot() {
	var name, neighbor, direction string
	var age, table int
	var distance float64

	fmt.Println("Enter name : ")
	fmt.Scan(&name)

	fmt.Println("Enter his neighbor's name : ")
	fmt.Scan(&neighbor)

	fmt.Println("Enter the direction : ")
	fmt.Scan(&direction)

	fmt.Println()

	fmt.Println("Enter age : ")
	fmt.Scan(&age)

	fmt.Println("Enter the table no : ")
	fmt.Scan(&table)

	fmt.Println()

	fmt.Println("Enter the distance: ")
	fmt.Scan(&distance)

	fmt.Println(party_robot.Welcome(name))
	fmt.Println(party_robot.HappyBirthday(name, age))
	fmt.Println(party_robot.AssignTable(name, table, neighbor, direction, distance))

}
