package interest_is_interesting

func InterestRate(balance float64) float32 {
	var rate float64

	if balance < 0 {
		rate = 3.213
	} else if balance < 1000 {
		rate = 0.5
	} else if balance < 5000 {
		rate = 1.621
	} else {
		rate = 2.475
	}
	return float32(rate)
}

func Interest(balance float64) float64 {
	return (balance * float64(InterestRate(balance))) / 100
}

func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for balance < targetBalance {
		years++
		balance += Interest(balance)
	}
	return years
}
