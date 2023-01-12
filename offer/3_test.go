package offer

import (
	"algorithm/utils"
	"math/rand"
	"testing"
)

func TestCountBits(t *testing.T) {

	const limit = 10000
	for i := 0; i < 10000; i++ {
		N := rand.Intn(limit)
		a := CountBits1(N)
		b := CountBits2(N)
		c := CountBits3(N)

		if !utils.IntSlicesAreEqual(a, b) || !utils.IntSlicesAreEqual(b, c) {
			t.Fatalf("Unequal N: %d!!!", N)
		}
	}
}

func benchmarkCountBits(b *testing.B, countBits func(int) []int) {
	for n := 0; n < b.N; n++ {
		countBits(n)
	}
}

func BenchmarkCountBits1(b *testing.B) { benchmarkCountBits(b, CountBits1) }
func BenchmarkCountBits2(b *testing.B) { benchmarkCountBits(b, CountBits2) }
func BenchmarkCountBits3(b *testing.B) { benchmarkCountBits(b, CountBits3) }
