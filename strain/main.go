package strain

import "fmt"

func Strain() {
	list := Lists{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	fmt.Println(list.Keep(filterInt))
}
