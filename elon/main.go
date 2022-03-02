package elon

import "fmt"

func Elon() {
	car := NewCar(100, 10)
	car.Drive()
	fmt.Println(car.DisplayDistance())
	fmt.Println(car.DisplayBattery())
	fmt.Println(car.CanFinish(100))
	fmt.Println(car.CanFinish(200))
}
