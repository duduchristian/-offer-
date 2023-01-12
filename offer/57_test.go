package offer

import (
	"fmt"
	"testing"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	fmt.Println(ContainsNearbyAlmostDuplicateB([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(ContainsNearbyAlmostDuplicateB([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	fmt.Println(ContainsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(ContainsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
