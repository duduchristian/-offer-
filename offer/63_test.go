package offer

import "testing"

func TestReplaceWords(t *testing.T) {
	if ReplaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery") != "the cat was rat by the bat" {
		t.Fatal("Error!")
	}
}
