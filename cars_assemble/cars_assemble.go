package cars_assemble

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate) / 100
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate)*successRate/100) / 60
}

func CalculateCost(carsCount int) uint {
	const individualCost = 10000
	const comboCost = 95000
	if carsCount < 10 {
		return uint(carsCount) * individualCost
	}
	return uint(carsCount/10)*comboCost + uint(carsCount%10)*individualCost
}
