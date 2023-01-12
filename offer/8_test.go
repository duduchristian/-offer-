package offer

import (
	"algorithm/utils"
	"math/rand"
	"testing"
	"time"
)

func TestMinSubArrayLen(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cases := []*struct {
		nums []int
		k    int
		res  int
	}{
		{
			[]int{5, 1, 4, 3},
			7,
			2,
		},
		{
			[]int{7, 9, 9, 7},
			7,
			1,
		},
	}

	for i := 0; i < 10000; i++ {
		k := rand.Intn(1000) + 150
		nums := utils.GenRandomSlice(1000, 200)
		res := MinSubArrayLenBruteForce(k, nums)
		cases = append(cases, &struct {
			nums []int
			k    int
			res  int
		}{nums: nums, k: k, res: res})
	}

	for _, c := range cases {
		res := MinSubArrayLen(c.k, c.nums)
		if res != c.res {
			t.Fatalf("Result should be %d instead of %d!!!", c.res, res)
		}
	}
}

func benchmarkMinSubArrayLen(b *testing.B, minSubArrayLen func(int, []int) int) {
	length := 10000
	k := length * 20
	nums := utils.GenRandomSlice(length, 200)
	for i := 0; i < b.N; i++ {
		minSubArrayLen(k, nums)
	}
}

func BenchmarkMinSubArrayLen(b *testing.B) { benchmarkMinSubArrayLen(b, MinSubArrayLen) }
func BenchmarkMinSubArrayLenBruteForce(b *testing.B) {
	benchmarkMinSubArrayLen(b, MinSubArrayLenBruteForce)
}
