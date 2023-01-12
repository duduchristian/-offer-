package offer

import (
	"math/rand"
	"testing"
	"time"
)

func TestSingleNumber(t *testing.T) {
	var cases [][]int
	rand.Seed(time.Now().UnixNano())
	const limit = 10000
	for i := 0; i < 1000; i++ {
		cases = append(cases, getRandomSlice(limit))
	}

	for _, c := range cases {
		trueRes := SingleNumberBruteForce(c)
		testRes := SingleNumber(c)
		if trueRes != testRes {
			t.Fatalf("Result should be %d but not %d!!!", trueRes, testRes)
		}
	}
}

func getRandomSlice(limit int) []int {
	num1 := rand.Int()
	c := []int{num1}
	for j := 0; j < limit; j++ {
		num := rand.Int()
		for k := 0; k < 3; k++ {
			c = append(c, num)
		}
	}
	return c
}

func benchmarkSingleNumber(b *testing.B, singleNumber func([]int) int) {
	c := getRandomSlice(10000)
	for i := 0; i < b.N; i++ {
		singleNumber(c)
	}
}

func BenchmarkSingleNumber(b *testing.B)           { benchmarkSingleNumber(b, SingleNumber) }
func BenchmarkSingleNumberBruteForce(b *testing.B) { benchmarkSingleNumber(b, SingleNumberBruteForce) }
