package dna

import "fmt"

func Nucleotide() {
	var nucleotide DNA = "ACGTTGCA"
	nucleotides, err := nucleotide.Counts()
	if err != nil {
		panic(err)
	}
	fmt.Println(nucleotides)
}
