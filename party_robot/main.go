package party_robot

import "fmt"

func Party_robot() {
	var name, neighbor, direction string
	var age, table int
	var distance float64

	fmt.Println("Enter name : ")
	fmt.Scan(&name)

	fmt.Println("Enter his neighbor's name : ")
	fmt.Scan(&neighbor)

	fmt.Println("Enter the direction : ")
	fmt.Scan(&direction)

	fmt.Println()

	fmt.Println("Enter age : ")
	fmt.Scan(&age)

	fmt.Println("Enter the table no : ")
	fmt.Scan(&table)

	fmt.Println()

	fmt.Println("Enter the distance: ")
	fmt.Scan(&distance)

	fmt.Println(Welcome(name))
	fmt.Println(HappyBirthday(name, age))
	fmt.Println(AssignTable(name, table, neighbor, direction, distance))
}
