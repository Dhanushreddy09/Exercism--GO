package triangle

import "testing"

func testKindFromSides(t *testing.T, a, b, c float64, expected Kind) {
	result := KindFromSides(a, b, c)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func benchmarkFromSides(b *testing.B, a, d, c float64) {
	for n := 0; n < b.N; n++ {
		KindFromSides(a, d, c)
	}
}

func TestKindFromSides(t *testing.T) { testKindFromSides(t, 1, 1, 1, Equ) }

func BenchmarkFromSides(b *testing.B) { benchmarkFromSides(b, 1, 3, 5) }
