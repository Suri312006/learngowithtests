package propbasedtests

import "testing"

func TestRomanNumberals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{
			"1 becomes I",
			1,
			"I",
		},
		{
			"2 becomes II",
			2,
			"II",
		},
		{
			"3 becomes III",
			3,
			"III",
		},
		{
			"4 becomes IV",
			4,
			"IV",
		},
		{
			"8 becomes VIII",
			8,
			"VIII",
		},
		{
			"9 becomes IX",
			9,
			"IX",
		},

		{
			"10 becomes X",
			10,
			"X",
		},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, wanted %q", got, test.Want)
			}

		})
	}
}
