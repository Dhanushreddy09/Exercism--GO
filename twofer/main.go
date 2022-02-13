package twofer

import "fmt"

func Twofer() {
	var name string

	fmt.Println("Enter name : ")
	fmt.Scanf("%s", &name)

	fmt.Println(ShareWith(name))
}
