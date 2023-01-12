package offer

//面试题43：在完全二叉树中添加节点

type CBTInserter struct {
	queue Queue
	root  TreeNode
}

func NewCBTInserter(root TreeNode) *CBTInserter {
	queue := NewQueue()
	queue.Offer(root)
	for queue.Peek().(TreeNode).left != nil && queue.Peek().(TreeNode).right != nil {
		node := queue.Poll().(TreeNode)
		queue.Offer(node.left)
		queue.Offer(node.right)
	}
	return &CBTInserter{
		queue: queue,
		root:  root,
	}
}

func (p *CBTInserter) Insert(v int) int {
	node := NewTreeNode(v)
	parent := p.queue.Peek().(TreeNode)
	if parent.left == nil {
		parent.left = node
	} else {
		parent.right = node

		p.queue.Poll()
		p.queue.Offer(parent.left)
		p.queue.Offer(parent.right)
	}
	return parent.val
}

func (p *CBTInserter) GetRoot() TreeNode {
	return p.root
}
