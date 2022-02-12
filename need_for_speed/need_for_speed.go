package need_for_speed

type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}
type Track struct {
	distance int
}

func NewCar(speed, batteryDrain int) Car {
	car := Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
	return car
}
func NewTrack(distance int) Track {
	track := Track{
		distance: distance,
	}
	return track
}
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}
	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}
func CanFinish(car Car, track Track) bool {
	distanceCovered := (car.battery / car.batteryDrain) * car.speed
	return distanceCovered >= track.distance

}
