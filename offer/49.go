package offer

//面试题49：从根节点到叶节点的路径数字之和

func SumNumbers(root TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	sumNumbersCore(root, 0, &ret)
	return ret
}

func sumNumbersCore(root TreeNode, cur int, sum *int) {
	curNum := cur*10 + root.val
	if root.left == nil && root.right == nil {
		*sum += curNum
		return
	}

	if root.left != nil {
		sumNumbersCore(root.left, curNum, sum)
	}

	if root.right != nil {
		sumNumbersCore(root.right, curNum, sum)
	}
}
