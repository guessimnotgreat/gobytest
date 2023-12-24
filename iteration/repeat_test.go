package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	expectedRepeatTime := 5
	repeated := Repeat("a", expectedRepeatTime)
	repeatedTime := len(repeated)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

	if expectedRepeatTime != repeatedTime {
		t.Errorf("expected '%d' but got '%d'", expectedRepeatTime, repeatedTime)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
