package romannumerals

import "testing"

func testToRomanNumeral(t *testing.T, number int, expected string) {
	actual, err := ToRomanNumeral(number)
	if err != nil {
		t.Errorf("ToRomanNumeral(%d) returned error: %s", number, err)
	} else if actual != expected {
		t.Errorf("ToRomanNumeral(%d) returned %s, expected %s", number, actual, expected)
	}
}

func benchmarkToRomanNumeral(b *testing.B, number int) {
	for i := 0; i < b.N; i++ {
		ToRomanNumeral(number)
	}
}

func TestToRomanNumeral(t *testing.T)        { testToRomanNumeral(t, 25, "XXV") }
func TestToRomanNumeral95(t *testing.T)      { testToRomanNumeral(t, 95, "XCV") }
func BenchmarkToRomanNumeral(b *testing.B)   { benchmarkToRomanNumeral(b, 25) }
func BenchmarkToRomanNumeral95(b *testing.B) { benchmarkToRomanNumeral(b, 95) }
func TestToRomanNumeral100(t *testing.T)     { testToRomanNumeral(t, 100, "C") }
