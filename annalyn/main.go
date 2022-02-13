package annalyn

import "fmt"

func Annalyn() {
	var knightIsAwake, archerIsAwake, prisonerIsAwake bool

	fmt.Println("Enter knightIsAwake : ")
	fmt.Scan(&knightIsAwake)

	fmt.Println("Enter archerIsAwake : ")
	fmt.Scan(&archerIsAwake)

	fmt.Println("Enter prisonerIsAwake : ")
	fmt.Scan(&prisonerIsAwake)

	fmt.Println("Can Fast Attack : ", CanFastAttack(knightIsAwake))
	fmt.Println("Can Spy : ", CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	fmt.Println("Can Signal Prisoner : ", CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
	fmt.Println("Can Free Prisoner : ", CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, false))
}
