package strain

import "fmt"

func Strain() {
	var list Ints
	list = append(list, 1, 2, 3, 4, 5)
	fmt.Println(list.Keep(filterInt))
}
