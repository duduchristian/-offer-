package offer

import "testing"

func TestAddBinary(t *testing.T) {
	cases := []*struct {
		a, b string
		res  string
	}{
		{"101", "101", "1010"},
		{"111", "111", "1110"},
		{"1000", "111", "1111"},
		{"1000000000", "1000000000", "10000000000"},
	}
	for _, c := range cases {
		res := AddBinary(c.a, c.b)
		if res != c.res {
			t.Fatalf("Result should be %s but not %s!!!", c.res, res)
		}
	}
}
