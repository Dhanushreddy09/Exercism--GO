package reverse_string

func Reverse(input string) string {
	inputRune := []rune(input)
	for i, j := 0, len(inputRune)-1; i < j; i, j = i+1, j-1 {
		inputRune[i], inputRune[j] = inputRune[j], inputRune[i]
	}
	return string(inputRune)
}
