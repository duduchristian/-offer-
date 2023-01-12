package offer

//面试题47：二叉树剪枝

func PruneTree(root TreeNode) TreeNode {
	if root == nil {
		return nil
	}
	root.left = PruneTree(root.left)
	root.right = PruneTree(root.right)
	if root.right == nil && root.left == nil && root.val == 0 {
		return nil
	}
	return root
}

func PruneTreeIterate(root TreeNode) TreeNode {
	stack := NewStack()
	var prev TreeNode
	cur := root

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

			if cur.right != nil && cur.right.left == nil &&
				cur.right.right == nil && cur.right.val == 0 {
				cur.right = nil
			}

			if cur.left != nil && cur.left.left == nil &&
				cur.left.right == nil && cur.left.val == 0 {
				cur.left = nil
			}

			prev = cur
			cur = nil
		}
	}

	if root.left == nil && root.right == nil && root.val == 0 {
		root = nil
	}

	return root
}
