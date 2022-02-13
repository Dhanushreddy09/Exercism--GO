package vehicle_purchase

import "fmt"

func Vehicle_purchase() {
	var kind, option1, option2 string

	fmt.Println("Enter kind of vehicle : ")
	fmt.Scan(&kind)

	fmt.Println("Enter option 1 : ")
	fmt.Scan(&option1)

	fmt.Println("Enter option2 : ")
	fmt.Scan(&option2)

	fmt.Println(NeedsLicense(kind))
	fmt.Println(ChooseVehicle(option1, option2))
	fmt.Println(CalculateResellPrice(10000.0, 5.0))
}
