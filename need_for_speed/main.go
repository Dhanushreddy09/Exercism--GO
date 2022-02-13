package need_for_speed

import "fmt"

func Need_for_speed() {
	var speed, batteryDrain, trackDistance int
	fmt.Println("Enter the speed of the vehicle : ")
	fmt.Scan(&speed)

	fmt.Println("Enter the battery drain : ")
	fmt.Scan(&batteryDrain)

	fmt.Println("Enter the track distance : ")
	fmt.Scan(&trackDistance)

	car := NewCar(speed, batteryDrain)
	track := NewTrack(trackDistance)

	fmt.Printf("Can finish the race : %v", CanFinish(car, track))
}
