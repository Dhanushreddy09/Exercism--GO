package elon

import "fmt"

func (car *Car) Drive() {
	if car.battery > car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {
	distanceCovered := (car.battery / car.batteryDrain) * car.speed
	return distanceCovered >= trackDistance
}
