package reverse_string

func Reverse(input string) string {
	var ans string
	inputRune := []rune(input)
	for i := len(inputRune) - 1; i >= 0; i-- {
		ans += string(inputRune[i])
	}
	return ans
}
