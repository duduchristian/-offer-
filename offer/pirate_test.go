package offer

import "testing"

func BenchmarkGetBestPlans(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetBestPlans(15, 100)
	}
}
