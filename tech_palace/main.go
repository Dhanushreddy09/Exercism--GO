package tech_palace

import "fmt"

func Tech_palace() {
	var customer string
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

	fmt.Println("Enter customer : ")
	fmt.Scan(&customer)

	fmt.Println(WelcomeMessage(customer))
	fmt.Println(AddBorder(WelcomeMessage(customer), 5))
	fmt.Println(CleanupMessage(message))
}
