package elon

import "fmt"

func Elon() {
	car := NewCar(100, 10)
	car.Drives()
	fmt.Println(car.DisplayDistances())
	fmt.Println(car.DisplayBatterys())
	fmt.Println(car.CanFinishorNot(100))
	fmt.Println(car.CanFinishorNot(200))
}
