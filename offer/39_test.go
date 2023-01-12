package offer

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	fmt.Println(LargestRectangleArea([]int{1, 1000, 1, 1}))
	fmt.Println(LargestRectangleArea([]int{3, 2, 5, 4, 6, 1, 4, 2}))
}
