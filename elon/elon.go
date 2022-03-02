package elon

import "fmt"

func (car *Car) Drives() {
	if car.battery > car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car Car) DisplayDistances() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

func (car Car) DisplayBatterys() string {
	return fmt.Sprintf("Battery at %v%%", car.battery)
}

func (car Car) CanFinishorNot(trackDistance int) bool {
	distanceCovered := (car.battery / car.batteryDrain) * car.speed
	return distanceCovered >= trackDistance
}
