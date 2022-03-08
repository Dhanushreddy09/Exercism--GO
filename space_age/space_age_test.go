package space_age

import "testing"

func testAge(t *testing.T, planet Planet, seconds float64, expected float64) {
	actual := Age(seconds, planet)
	if actual != expected {
		t.Errorf("Age(%f) = %f, want %f", seconds, actual, expected)
	}
}

func benchmarkAge(b *testing.B, planet Planet, seconds float64) {
	for i := 0; i < b.N; i++ {
		Age(seconds, planet)
	}
}

func TestAge(t *testing.T)          { testAge(t, "Mars", 1000000000, 41.88) }
func BenchmarkMars(b *testing.B)    { benchmarkAge(b, "Mars", 1000000000) }
func TestAge2(t *testing.T)         { testAge(t, "Jupiter", 1000000000, 2.41) }
func BenchmarkJupiter(b *testing.B) { benchmarkAge(b, "Jupiter", 1000000000) }
func TestAge3(t *testing.T)         { testAge(t, "Saturn", 2134835688, 95.06) }
func BenchmarkSaturn(b *testing.B)  { benchmarkAge(b, "Saturn", 2134835688) }
func TestAge4(t *testing.T)         { testAge(t, "Neptune", 1821023256, 30.12) }
func BenchmarkNeptune(b *testing.B) { benchmarkAge(b, "Neptune", 1821023256) }
