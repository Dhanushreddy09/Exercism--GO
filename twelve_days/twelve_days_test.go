package twelve_days

import "testing"

func testVerse(t *testing.T, number int, expected string) {
	actual := Verse(number)
	if actual != expected {
		t.Errorf("Verse(%d) = %q, want %q", number, actual, expected)
	}
}

func benchmarkVerse(b *testing.B, number int) {
	for i := 0; i < b.N; i++ {
		Verse(number)
	}
}

func benchmarkSong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Song()
	}
}

func TestVerse(t *testing.T) {
	testVerse(t, 1, "On the first day of Christmas my true love gave to me: a Partridge in a Pear Tree.")
}

func BenchmarkVerse(b *testing.B) { benchmarkVerse(b, 1) }
func BenchmarkSong(b *testing.B)  { benchmarkSong(b) }
