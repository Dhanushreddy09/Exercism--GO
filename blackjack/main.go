package blackjack

import "fmt"

func Blackjack() {
	var card1, card2 string
	var dealerScore, handScore int

	fmt.Println("Enter card1 : ")
	fmt.Scan(&card1)

	fmt.Println("Enter card2 : ")
	fmt.Scan(&card2)

	fmt.Println("Enter dealer score : ")
	fmt.Scan(&dealerScore)

	fmt.Println("Enter handScore : ")
	fmt.Scan(&handScore)

	fmt.Println(ParseCard(card1))
	blackjackVal := IsBlackjack(card1, card2)
	fmt.Println(LargeHand(blackjackVal, handScore))
}
