package triangle

type Kind string

const (
	NaT Kind = "Not a triangle"
	Equ Kind = "Equilateral Triangle"
	Iso Kind = "Isosceles Triangle"
	Sca Kind = "Scalene Triangle"
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a <= 0 || b <= 0 || c <= 0 || !determineTriangle(a, b, c) {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}

func determineTriangle(a, b, c float64) bool {
	if a+b <= c || b+c <= a || c+a <= b {
		return false
	}
	return true
}
