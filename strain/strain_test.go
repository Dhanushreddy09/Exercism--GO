package strain

import (
	"reflect"
	"testing"
)

func testInts(t *testing.T, input Ints, filter func(int) bool, expected Ints) {
	actual := input.Keep(filter)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Keep(%v) = %v, expected %v", input, actual, expected)
	}
}

func testLists(t *testing.T, input Lists, filter func(int) bool, expected Lists) {
	actual := input.Keep(filter)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Keep(%v) = %v, expected %v", input, actual, expected)
	}
}

func testStrings(t *testing.T, input Strings, filter func(string) bool, expected Strings) {
	actual := input.Keep(filter)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Keep(%v) = %v, expected %v", input, actual, expected)
	}
}

func benchmarkInts(b *testing.B, input Ints) {
	for i := 0; i < b.N; i++ {
		input.Keep(filterInt)
	}
}

func benchmarkLists(b *testing.B, input Lists) {
	for i := 0; i < b.N; i++ {
		input.Keep(filterInt)
	}
}

func benchmarkStrings(b *testing.B, input Strings) {
	for i := 0; i < b.N; i++ {
		input.Keep(filterString)
	}
}

func TestKeep(t *testing.T) {
	ints := Ints{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lists := Lists{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	strings := Strings{
		"hello",
		"world",
		"this",
		"is",
		"a",
		"testament",
	}

	testInts(t, ints, filterInt, Ints{2, 4, 6, 8, 10})
	testLists(t, lists, filterInt, Lists{{2}, {4, 6}, {8}, {10, 12}})
	testStrings(t, strings, filterString, Strings{"testament"})
}
func BenchmarkKeepInts(b *testing.B) { benchmarkInts(b, Ints{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) }
func BenchmarkDiscardstrings(b *testing.B) {
	benchmarkStrings(b, Strings{"Hello", "world", "this", "is", "a", "testament"})
}
func BenchmarkKeepLists(b *testing.B) {
	benchmarkLists(b, Lists{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	})
}
