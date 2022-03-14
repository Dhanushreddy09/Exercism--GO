package anagram

import "testing"

func testAnagram(t *testing.T, subject string, candidates []string, expected []string) {
	result := Detect(subject, candidates)
	if len(result) != len(expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	}
}

func benchmarkAnagram(b *testing.B, subject string, candidates []string) {
	for n := 0; n < b.N; n++ {
		Detect(subject, candidates)
	}
}

func TestAnagram(t *testing.T) {
	testAnagram(t, "Orchestra", []string{"cashregister", "Carthorse", "radishes"}, []string{"Carthorse"})
}

func BenchmarkAnagram(b *testing.B) {
	benchmarkAnagram(b, "Orchestra", []string{"cashregister", "Carthorse", "radishes"})
}
