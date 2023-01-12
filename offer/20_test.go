package offer

import "testing"

func TestCountSubstrings(t *testing.T) {
	cases := []*struct {
		s     string
		count int
	}{
		{
			"abca", 4,
		},
		{
			"aba", 4,
		},
		{
			"ab", 2,
		},
		{
			"aaba", 6,
		},
		{
			"baaba", 8,
		},
	}

	for _, c := range cases {
		testRes := CountSubstrings(c.s)
		if testRes != c.count {
			t.Fatalf("%s is %d", c.s, c.count)
		}
	}
}
