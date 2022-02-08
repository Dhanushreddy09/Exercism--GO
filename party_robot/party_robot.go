package party_robot

import "fmt"

func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %v! You are now %v years old!", name, age)
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var msg string
	msg += fmt.Sprintf("Welcome to my party, %v!", name)
	msg += fmt.Sprintf("\nYou have been assigned to table %03d. Your table is %v, exactly %.1f meters from here.", table, direction, distance)
	msg += fmt.Sprintf("\nYou will be sitting next to %v.", neighbor)
	return msg
}
