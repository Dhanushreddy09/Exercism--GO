package rna_transcription

func ToRNA(dna string) string {
	dnaRune := []rune(dna)
	for i, nucleo := range dnaRune {
		switch nucleo {
		case 'G':
			dnaRune[i] = 'C'
		case 'C':
			dnaRune[i] = 'G'
		case 'T':
			dnaRune[i] = 'A'
		case 'A':
			dnaRune[i] = 'U'
		}
	}
	return string(dnaRune)
}
