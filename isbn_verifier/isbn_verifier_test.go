package isbn_verifier

import "testing"

func testIsValidISBN(t *testing.T, isbn string, expected bool) {
	actual := IsValidISBN(isbn)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func benchmarkIsValidISBN(b *testing.B, isbn string) {
	for i := 0; i < b.N; i++ {
		IsValidISBN(isbn)
	}
}

func TestIsValidISBN(t *testing.T)      { testIsValidISBN(t, "359821507X", true) }
func BenchmarkIsValidISBN(b *testing.B) { benchmarkIsValidISBN(b, "359821507X") }
