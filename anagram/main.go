package anagram

import "fmt"

func Anagram() {
	subject := "abc"
	candidates := []string{"abc", "cba"}
	fmt.Println(Detect(subject, candidates))
}
