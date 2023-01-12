package offer

//面试题50：向下的路径节点值之和

func PathSum(root TreeNode, target int) int {
	return pathSumCore(root, target, 0, map[int]int{})
}

func pathSumCore(root TreeNode, target, sum int, mem map[int]int) int {
	if root == nil {
		return 0
	}
	mem[sum+root.val]++
	res := mem[target-root.val] +
		pathSumCore(root.left, target, sum+root.val, mem) +
		pathSumCore(root.right, target, sum+root.val, mem)
	mem[sum+root.val]--
	return res
}
