package offer

//面试题52：展平二叉搜索树

func IncreasingBST(root TreeNode) TreeNode {
	stack := NewStack()
	cur := root
	dummy := NewTreeNode(0)
	prev := dummy

	for !stack.Empty() && cur != nil {
		for cur != nil {
			stack.Push(cur)
			cur = cur.left
		}

		cur = stack.Pop().(TreeNode)
		prev.right = cur
		prev = cur

		cur.left = nil
		cur = cur.right
	}

	return dummy.right
}
