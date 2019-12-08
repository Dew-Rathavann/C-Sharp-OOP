package Hello

import "testing"

func TestFind(t *testing.T) {
	t.Run("Hello world", func(t *testing.T) {
		want := true
		got := Find("Hello world")
		if got != want {
			t.Errorf("\ngot %v want %v", got, want)
		}
	})
	t.Run("Hello Solar System", func(t *testing.T) {
		want := false
		got := Find("Hello Solar System")
		if got != want {
			t.Errorf("\ngot %v want %v", got, want)
		}
	})
	t.Run("Hello Lunar", func(t *testing.T) {
		want := false
		got := Find("Hello Lunar")
		if got != want {
			t.Errorf("\ngot %v want %v", got, want)
		}
	})
	t.Run("Hello we are the CS", func(t *testing.T) {
		want := true
		got := Find("Hello we are the CS")
		if got != want {
			t.Errorf("\ngot %v want %v", got, want)
		}
	})
}
