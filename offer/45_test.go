package offer

import (
	"fmt"
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {
	root := ConstructCBT([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(FindBottomLeftValue(root))
}
