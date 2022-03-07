package leap

import "testing"

func testIsLeapYear(t *testing.T, year int, expected bool) {
	if IsLeapYear(year) != expected {
		t.Fatalf("%v is a leap year: %v\n", year, expected)
	}
}

func benchmarkIsLeapYear(b *testing.B, year int) {
	for i := 0; i < b.N; i++ {
		IsLeapYear(year)
	}
}

func TestIsLeapYear(t *testing.T)      { testIsLeapYear(t, 2024, true) }
func BenchmarkIsLeapYear(b *testing.B) { benchmarkIsLeapYear(b, 2028) }
