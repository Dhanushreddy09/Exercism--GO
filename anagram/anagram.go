package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	alphabets := map[rune]int{}
	ans := []string{}
	contains := true

	for _, val := range subject {
		_, exists := alphabets[val]
		if !exists {
			alphabets[val] = 1
		} else {
			alphabets[val]++
		}
	}

	for _, candidate := range candidates {
		candidateName := strings.ToLower(candidate)
		if !strings.EqualFold(subject, candidate) {
			candidateAlpha := map[rune]int{}
			for _, candidateRune := range candidateName {
				_, exists := alphabets[candidateRune]
				if !exists || len(subject) != len(candidateName) {
					contains = !contains
					break
				}
				_, runeExists := candidateAlpha[candidateRune]
				if !runeExists {
					candidateAlpha[candidateRune] = 1
				} else {
					candidateAlpha[candidateRune]++
				}
				if candidateAlpha[candidateRune] > alphabets[candidateRune] {
					contains = !contains
				}
			}
			if contains {
				ans = append(ans, candidate)
			} else {
				contains = !contains
			}
		}
	}
	return ans
}
