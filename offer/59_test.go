package offer

import (
	"algorithm/utils"
	"sort"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	k := NewKthLargest(3, []int{4, 5, 8, 2})
	if k.Add(3) != 4 {
		t.Fatal("Error!")
	}

	s1 := utils.GenRandomSlice(1000, 1000)
	s2 := utils.GenRandomSlice(100, 1000)

	k = NewKthLargest(20, s1)
	for _, num := range s2 {
		s1 = append(s1, num)
		sort.Slice(s1, func(i, j int) bool {
			return s1[i] > s1[j]
		})
		if k.Add(num) != s1[19] {
			t.Fatal("Error!!")
		}
	}
}
