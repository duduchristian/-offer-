package offer

import (
	"fmt"
	"testing"
)

func TestPruneTree(t *testing.T) {
	root1 := ConstructCBT([]int{1, 0, 0, 0, 0, 0, 1})
	root2 := ConstructCBT([]int{1, 0, 0, 0, 0, 0, 1, 0, 0})

	PruneTree(root1)
	PruneTreeIterate(root2)

	fmt.Println(IsEqual(root1, root2))
}
