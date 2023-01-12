package offer

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("tasty")
	trie.Insert("great")
	trie.Insert("what's up")
	if !trie.Search("tasty") {
		t.Fatal("tasty should be true")
	}
	if !trie.Search("great") {
		t.Fatal("great should be true")
	}
	if !trie.Search("what's up") {
		t.Fatal("what's up should be true")
	}
	if trie.Search("abc") {
		t.Fatal("abc should be false")
	}
	if !trie.StartsWith("what") {
		t.Fatal("what should be true")
	}
	if trie.StartsWith("what''s") {
		t.Fatal("what''s should be false")
	}
}
