package offer

type MagicDictionary struct {
	*Trie
}

func NewMagicDictionary(dict []string) *MagicDictionary {
	ret := &MagicDictionary{
		NewTrie(),
	}
	for _, word := range dict {
		ret.Insert(word)
	}
	return ret
}

func (d *MagicDictionary) Search(word string) bool {
	return d.dfs(d.root, []rune(word), 0, 0)
}

func (d *MagicDictionary) dfs(root TrieNode, rs []rune, i, edit int) bool {
	if root == nil {
		return false
	}
	if root.isWord && i == len(rs) && edit == 1 {
		return true
	}
	if i < len(rs) && edit <= 1 {
		for r, child := range root.children {
			next := edit
			if r != rs[i] {
				next++
			}
			if d.dfs(child, rs, i+1, next) {
				return true
			}
		}
	}
	return false
}
