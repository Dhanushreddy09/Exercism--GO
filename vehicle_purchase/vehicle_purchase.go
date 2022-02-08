package vehicle_purchase

func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}
func ChooseVehicle(option1, option2 string) string {
	helperString := " is clearly the better choice."

	if option1 < option2 {
		return option1 + helperString
	}
	return option2 + helperString
}
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return (originalPrice * 80) / 100
	} else if age >= 3 && age < 10 {
		return (originalPrice * 70) / 100
	}
	return (originalPrice * 50) / 100
}
