package dna

import "errors"

type Histogram map[rune]int
type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, val := range d {
		_, exists := h[val]
		if !exists {
			return h, errors.New("nucleotides must be A, C, G, or T")
		}
		h[val]++
	}
	return h, nil
}
