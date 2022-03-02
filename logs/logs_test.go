package logs

import "testing"

func testApplication(log, expectedApp string, t *testing.T) {
	actual := Application(log)
	if actual != expectedApp {
		t.Errorf("Application(%q) should be %q but was %q", log, expectedApp, actual)
	}
}

func benchmarkApplication(log string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Application(log)
	}
}

func TestApplication(t *testing.T) { testApplication("This is a search icon 🔍.", "search", t) }

func BenchmarkApplication(b *testing.B)  { benchmarkApplication("This is a search icon 🔍.", b) }
func BenchmarkApplication2(b *testing.B) { benchmarkApplication("This is a recommended icon ❗.", b) }
