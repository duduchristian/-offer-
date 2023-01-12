package offer

import (
	"testing"
)

func TestMinimumLengthEncoding(t *testing.T) {
	if MinimumLengthEncoding([]string{"time", "me", "bell"}) != 10 {
		t.Fatal("Error!")
	}
}
