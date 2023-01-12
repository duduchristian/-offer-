package offer

import "testing"

func TestValidPalindrome(t *testing.T) {
	cases := []*struct {
		s            string
		isPalindrome bool
	}{
		{
			"abca", true,
		},
		{
			"aba", true,
		},
		{
			"ab", true,
		},
		{
			"abc", false,
		},
		{
			"aaaaaaba", true,
		},
	}

	for _, c := range cases {
		testRes := ValidPalindrome(c.s)
		if testRes != c.isPalindrome {
			t.Fatalf("%s is %t", c.s, c.isPalindrome)
		}
	}
}
