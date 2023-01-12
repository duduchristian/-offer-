package offer

import (
	"fmt"
	"testing"
)

func TestTraversal(t *testing.T) {
	root := ConstructCBT([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(InorderTraversal(root))
	fmt.Println(PreorderTraversal(root))
	fmt.Println(PostorderTraversal(root))
}
