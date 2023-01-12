package offer

func InorderTraversal(root TreeNode) []int {
	stack := NewStack()
	cur := root
	var ret []int
	for !stack.Empty() || cur != nil {
		for cur != nil {
			stack.Push(cur)
			cur = cur.left
		}
		node := stack.Pop().(TreeNode)
		ret = append(ret, node.val)
		cur = node.right
	}
	return ret
}

func PreorderTraversal(root TreeNode) []int {
	stack := NewStack()
	cur := root
	var ret []int
	for !stack.Empty() || cur != nil {
		for cur != nil {
			ret = append(ret, cur.val)
			stack.Push(cur)
			cur = cur.left
		}
		node := stack.Pop().(TreeNode)
		cur = node.right
	}
	return ret
}

func PostorderTraversal(root TreeNode) []int {
	stack := NewStack()
	var prev TreeNode
	cur := root
	var ret []int
	for !stack.Empty() || cur != nil {
		for cur != nil {
			stack.Push(cur)
			cur = cur.left
		}

		cur = stack.Peek().(TreeNode)
		if cur.right != nil && cur.right != prev {
			cur = cur.right
		} else {
			stack.Pop()
			ret = append(ret, cur.val)
			prev = cur
			cur = nil
		}
	}
	return ret
}
