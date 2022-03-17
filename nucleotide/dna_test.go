package dna

import (
	"reflect"
	"testing"
)

func testCounts(t *testing.T, dna DNA, expected Histogram) {
	actual, err := dna.Counts()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func benchmarkCounts(dna DNA, b *testing.B) {
	for i := 0; i < b.N; i++ {
		dna.Counts()
	}
}

func TestCounts(t *testing.T) { testCounts(t, "ACGT", Histogram{'A': 1, 'C': 1, 'G': 1, 'T': 1}) }

func BenchmarkCounts(b *testing.B) { benchmarkCounts("AGTCTGAC", b) }
