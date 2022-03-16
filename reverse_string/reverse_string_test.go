package reverse_string

import "testing"

func testReverse(t *testing.T, input string, expected string) {
	actual := Reverse(input)
	if actual != expected {
		t.Errorf("Reverse(%q) == %q, expected %q", input, actual, expected)
	}
}

func benchmarkReverse(b *testing.B, input string) {
	for i := 0; i < b.N; i++ {
		Reverse(input)
	}
}

func TestReverse(t *testing.T) {
	testReverse(t, "Hello, World!", "!dlroW ,olleH")
}

func BenchmarkReverse(b *testing.B) { benchmarkReverse(b, "Hello, World!") }
