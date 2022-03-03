package census

import "testing"

func testHasRequiredInfo(t *testing.T, r *Resident, expected bool) {
	actual := r.HasRequiredInfo()
	if actual != expected {
		t.Errorf("Expected %t, got %t", expected, actual)
	}
}

func benchmarkHasRequiredInfo(b *testing.B, r *Resident) {
	for i := 0; i < b.N; i++ {
		r.HasRequiredInfo()
	}
}

func TestHasRequiredInfo(t *testing.T) {
	testHasRequiredInfo(t, NewResident("Dhanush", 20, map[string]string{"street": "oxford street"}), true)
}

func BenchmarkHasRequiredInfo(b *testing.B) {
	benchmarkHasRequiredInfo(b, NewResident("Dhanush", 20, map[string]string{"street": "oxford street"}))
}
