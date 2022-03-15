package triangle

import "fmt"

func Triangle() {
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)
	fmt.Println(KindFromSides(a, b, c))
}
