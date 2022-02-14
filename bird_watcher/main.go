package bird_watcher

import "fmt"

func Bird_watcher() {
	birdsPerDay := []int{5, 8, 9, 3, 2, 1, 4, 7, 6}
	fmt.Println("Total birds in the first week:", BirdsInWeek(birdsPerDay, 1))

}
