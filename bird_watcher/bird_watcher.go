package bird_watcher

func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, value := range birdsPerDay {
		sum += value
	}
	return sum
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	index := (week * 7) - 7
	for i := index; i < index+7; i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
