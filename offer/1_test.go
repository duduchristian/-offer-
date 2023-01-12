package offer

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestDivide(t *testing.T) {
	cases := []*struct {
		dividend int
		divisor  int
		result   int
	}{
		{
			math.MinInt,
			-1,
			math.MaxInt,
		},
		{
			math.MinInt,
			10086,
			math.MinInt / 10086,
		},
		{
			0,
			1,
			0,
		},
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000000; i++ {
		divisor := rand.Int()
		if divisor == 0 {
			i--
			continue
		}
		dividend := rand.Int()
		if rand.Float64() < 0.5 {
			dividend *= -1
		}
		if rand.Float64() < 0.5 {
			divisor *= -1
		}

		cases = append(cases, &struct {
			dividend int
			divisor  int
			result   int
		}{dividend: dividend, divisor: divisor, result: dividend / divisor})
	}

	for _, c := range cases {
		testRes := Divide(c.dividend, c.divisor)
		if c.result != testRes {
			t.Fatalf("Result should be %d but not %d", c.result, testRes)
		}
	}
}
