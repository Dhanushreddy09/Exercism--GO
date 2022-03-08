package space_age

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var ans float64
	earthYears := 31557600.00
	switch planet {
	case "Earth":
		ans = seconds / earthYears
	case "Mercury":
		ans = seconds / (earthYears * 0.2408467)
	case "Venus":
		ans = seconds / (earthYears * 0.61519726)
	case "Mars":
		ans = seconds / (earthYears * 1.8808158)
	case "Jupiter":
		ans = seconds / (earthYears * 11.862615)
	case "Saturn":
		ans = seconds / (earthYears * 29.447498)
	case "Uranus":
		ans = seconds / (earthYears * 84.016846)
	case "Neptune":
		ans = seconds / (earthYears * 164.79132)
	}
	return ans
}
