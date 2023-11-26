package loops

import "testing"

func TestRepeatDefaultTimes(t *testing.T) {
	loops := Repeat("a", 0)
	expected := "aaaaa"

	if loops != expected {
		t.Errorf("expected '%s' but returned '%s'", expected, loops)
	}
}

func TestRepeatWithNumber(t *testing.T) {
	loops := Repeat("a", 7)
	expected := "aaaaaaa"

	if loops != expected {
		t.Errorf("expected '%s' but returned '%s'", expected, loops)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
