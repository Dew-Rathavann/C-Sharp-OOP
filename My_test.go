package My

import (
	"testing"
)

func TestSum(t *testing.T) {
	want := 1
	got := Sum(1, 1)
	if got != want {
		t.Errorf("wanted %v but got %v instead\n", want, got)
	}
	want = 1
	got = Sum(2, 1)
	if got != want {
		t.Errorf("wanted %v but got %v instead\n", want, got)
	}
	want = 1
	got = Sum(2, 4)
	if got != want {
		t.Errorf("wanted %v but got %v instead\n", want, got)
	}
}
