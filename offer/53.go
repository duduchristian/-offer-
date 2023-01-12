package offer

//面试题53：二叉搜索树的下一个节点

func InorderSuccessor(root, p TreeNode) (res TreeNode) {
	cur := root
	for cur != nil {
		if cur.val > p.val {
			res = cur
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return
}
