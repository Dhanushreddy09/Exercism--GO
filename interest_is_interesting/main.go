package interest_is_interesting

import "fmt"

func Interest_is_interesting() {
	var balance, targetBalance float64

	fmt.Println("Enter balance : ")
	fmt.Scan(&balance)

	fmt.Println("Enter targetBalance : ")
	fmt.Scan(&targetBalance)

	fmt.Println(AnnualBalanceUpdate(balance))
	fmt.Println(YearsBeforeDesiredBalance(balance, targetBalance))
}
