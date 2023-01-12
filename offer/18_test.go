package offer

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []*struct {
		s            string
		isPalindrome bool
	}{
		{
			"Was it a cat I saw？",
			true,
		},
		{
			"race a car",
			false,
		},
	}

	for _, c := range cases {
		testRes := IsPalindrome(c.s)
		if testRes != c.isPalindrome {
			t.Fatalf("%s is %t", c.s, c.isPalindrome)
		}
	}
}
