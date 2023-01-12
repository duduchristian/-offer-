package offer

//面试题45：二叉树最低层最左边的值

func FindBottomLeftValue(root TreeNode) int {
	queue1 := NewQueue()
	queue2 := NewQueue()
	ret := root.val
	queue1.Offer(root)

	for !queue1.Empty() {
		node := queue1.Poll().(TreeNode)
		if node.left != nil {
			queue2.Offer(node.left)
		}
		if node.right != nil {
			queue2.Offer(node.right)
		}

		if queue1.Empty() {
			if !queue2.Empty() {
				ret = queue2.Peek().(TreeNode).val
			}
			queue1, queue2 = queue2, queue1
		}
	}
	return ret
}
