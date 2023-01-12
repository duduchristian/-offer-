package offer

type BSTIterator struct {
	cur   TreeNode
	stack Stack
}

func NewBSTIterator(root TreeNode) *BSTIterator {
	return &BSTIterator{
		cur:   root,
		stack: NewStack(),
	}
}

func (p *BSTIterator) HasNext() bool {
	return p.cur != nil || !p.stack.Empty()
}

func (p *BSTIterator) Next() int {
	for p.cur != nil {
		p.stack.Push(p.cur)
		p.cur = p.cur.left
	}

	p.cur = p.stack.Pop().(TreeNode)
	ret := p.cur.val
	p.cur = p.cur.right
	return ret
}
