package offer

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	cases := []*struct {
		s1  string
		s2  string
		res bool
	}{
		{
			"ca",
			"dgcaf",
			true,
		},
		{
			"afc",
			"dgcaf",
			true,
		},
		{
			"ab",
			"dgcaf",
			false,
		},
		{
			"",
			"",
			true,
		},
		{
			"",
			"ahhtoaht",
			true,
		},
		{
			"ahtohatl",
			"ah",
			false,
		},
		{
			"stop",
			"pots",
			true,
		},
	}

	for _, c := range cases {
		if c.res != CheckInclusion(c.s1, c.s2) {
			t.Fatalf("s1 %s s2 %s should be %t", c.s1, c.s2, c.res)
		}
	}
}
