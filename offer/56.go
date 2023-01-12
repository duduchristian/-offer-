package offer

type BSTIteratorReverted struct {
	cur   TreeNode
	stack Stack
}

func NewBSTIteratorReverted(root TreeNode) *BSTIteratorReverted {
	return &BSTIteratorReverted{
		cur:   root,
		stack: NewStack(),
	}
}

func (p *BSTIteratorReverted) HasNext() bool {
	return p.cur != nil || !p.stack.Empty()
}

func (p *BSTIteratorReverted) Next() int {
	for p.cur != nil {
		p.stack.Push(p.cur)
		p.cur = p.cur.right
	}

	p.cur = p.stack.Pop().(TreeNode)
	ret := p.cur.val
	p.cur = p.cur.left
	return ret
}

func FindTarget(root TreeNode, k int) bool {
	if root == nil {
		return false
	}

	left := NewBSTIterator(root)
	right := NewBSTIteratorReverted(root)

	next := left.Next()
	prev := right.Next()
	for next != prev {
		if next+prev == k {
			return true
		}
		if next+prev < k {
			next = left.Next()
		} else {
			prev = right.Next()
		}
	}
	return false
}
