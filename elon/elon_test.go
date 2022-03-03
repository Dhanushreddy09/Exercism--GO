package elon

import "testing"

func testDrives(t *testing.T, car *Car, expectedCar *Car) {
	car.Drives()
	if car.distance != expectedCar.distance {
		t.Errorf("Expected distance %v, got %v", expectedCar.distance, car.distance)
	}
	if car.battery != expectedCar.battery {
		t.Errorf("Expected battery %v, got %v", expectedCar.battery, car.battery)
	}
}

func benchmarkDrives(car *Car, b *testing.B) {
	for n := 0; n < b.N; n++ {
		car.Drives()
	}
}

func TestDrives(t *testing.T) {
	car := NewCar(5, 2)
	expectedCar := &Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
	testDrives(t, car, expectedCar)
}

func BenchmarkDrives(b *testing.B)   { benchmarkDrives(NewCar(5, 2), b) }
func BenchmarkDrives10(b *testing.B) { benchmarkDrives(NewCar(10, 8), b) }
