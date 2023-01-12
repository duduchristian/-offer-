package offer

import (
	"testing"
)

func TestMapSum(t *testing.T) {
	mapSum := NewMapSum()
	mapSum.Insert("happy", 3)
	mapSum.Insert("happen", 2)

	if mapSum.Sum("hap") != 5 {
		t.Fatal("Error!")
	}

	mapSum.Insert("h", 1)
	if mapSum.Sum("h") != 6 {
		t.Fatal("Error!!")
	}

	if mapSum.Sum("ha") != 5 {
		t.Fatal("Error!!!")
	}
}
