package offer

//面试题54：所有大于或等于节点的值之和

func ConvertBST(root TreeNode) TreeNode {
	stack := NewStack()
	cur := root
	prevSum := 0

	for !stack.Empty() || cur != nil {
		for cur != nil {
			stack.Push(cur)
			cur = cur.right
		}

		cur = stack.Pop().(TreeNode)
		prevSum += cur.val
		cur.val = prevSum
		cur = cur.left
	}
	return root
}
