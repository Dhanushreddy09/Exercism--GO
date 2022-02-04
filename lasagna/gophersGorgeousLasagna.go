package lasagna

const OvenTime int = 40
const layerPrepTime = 2

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * layerPrepTime
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
