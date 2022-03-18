package rna_transcription

import "testing"

func testToRna(t *testing.T, dna string, expected string) {
	actual := ToRNA(dna)
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func benchmarkToRna(b *testing.B, dna string) {
	for i := 0; i < b.N; i++ {
		ToRNA(dna)
	}
}

func TestToRna(t *testing.T)      { testToRna(t, "GATGGAACTTGACTACGTAAATT", "CUACCUUGAACUGAUGCAUUUAA") }
func BenchmarkToRna(b *testing.B) { benchmarkToRna(b, "GATGGAACTTGACTACGTAAATT") }
