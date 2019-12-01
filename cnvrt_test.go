package Convert

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Descripption string
		Arabic       int
		Want         string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConverttoRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
