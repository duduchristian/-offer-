package offer

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	root1 := ConstructCBT([]int{1, 0, 0, 0, 0, 0, 1})
	root2 := Deserialize(Serialize(root1))

	if !IsEqual(root1, root2) {
		t.Fatal("Error!")
	}
}
