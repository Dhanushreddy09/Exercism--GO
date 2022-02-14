package lasagna_master

import "fmt"

func Lasagna_master() {
	layers := []string{"sauce", "noodles", "cheese", "tomato", "mushroom", "onion", "pepper", "pepperoni", "sausage"}
	fmt.Println("Preparation time:", PreparationTime(layers, 2))
}
