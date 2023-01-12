package offer

import (
	"strings"
)

//面试题62：实现前缀树

type TrieNode *struct {
	children map[rune]TrieNode
	isWord   bool
}

func NewTrieNode() TrieNode {
	return &struct {
		children map[rune]TrieNode
		isWord   bool
	}{
		children: make(map[rune]TrieNode, 0),
	}
}

type Trie struct {
	root TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		if node.children[r] == nil {
			node.children[r] = NewTrieNode()
		}
		node = node.children[r]
	}
	node.isWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, r := range word {
		if child, ok := node.children[r]; ok {
			node = child
			continue
		}
		return false
	}
	return node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, r := range prefix {
		if child, ok := node.children[r]; ok {
			node = child
			continue
		}
		return false
	}
	return true
}

func (t *Trie) FindPrefix(word string) string {
	node := t.root
	builder := strings.Builder{}
	for _, r := range word {
		if node.isWord || node.children[r] == nil {
			break
		}
		builder.WriteRune(r)
		node = node.children[r]
	}
	if !node.isWord {
		return ""
	}
	return builder.String()
}
