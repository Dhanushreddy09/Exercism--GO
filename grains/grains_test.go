package grains

import "testing"

func benchmarkSquare(i int, b *testing.B) {
	for n := 0; i < b.N; n++ {
		Square(i)
	}
}

func testSquare(t *testing.T, i int, expected uint64) {
	actual, err := Square(i)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Square(%d) should be %d but was %d", i, expected, actual)
	}
}

func TestSquare(t *testing.T) {
	testSquare(t, 0, 1)
}

func BenchmarkSquare1(b *testing.B)  { benchmarkSquare(1, b) }
func BenchmarkSquare2(b *testing.B)  { benchmarkSquare(2, b) }
func BenchmarkSquare10(b *testing.B) { benchmarkSquare(10, b) }
func BenchmarkSquare20(b *testing.B) { benchmarkSquare(20, b) }
