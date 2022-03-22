package word_count

import (
	"reflect"
	"testing"
)

func testWordCount(t *testing.T, phrase string, expected Frequency) {
	actual := WordCount(phrase)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("WordCount(%q) = %q, want %q", phrase, actual, expected)
	}
}

func benchmarkWordCount(b *testing.B, phrase string) {
	for i := 0; i < b.N; i++ {
		WordCount(phrase)
	}
}

func TestWordCount(t *testing.T) {
	testWordCount(t, "car: carpet as java: javascript!!&@$%^&", Frequency{"car": 1, "carpet": 1, "as": 1, "java": 1, "javascript": 1})
}
func BenchmarkWordCount(b *testing.B) {
	benchmarkWordCount(b, "car: carpet as java: javascript!!&@$%^&")
}
