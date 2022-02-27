package luhn

func Valid(id string) bool {
	if len(id) <= 1 {
		return false
	}
	counter := 0
	sum := 0

	for i := len(id) - 1; i >= 0; i-- {
		char := id[i]
		if char == 32 {
			continue
		}
		if char < 48 || char > 57 {
			return false
		}
		digit := int(char - 48)
		counter++

		if counter%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
