package offer

import (
	"algorithm/utils"
)

//面试题51：节点值之和最大的路径

func MaxPathSum(root TreeNode) int {
	ret, _ := maxPathSumCore(root)
	return ret
}

func maxPathSumCore(root TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftMax, leftPath := maxPathSumCore(root.left)
	rightMax, rightPath := maxPathSumCore(root.right)
	leftPath = utils.Max(0, leftPath)
	rightPath = utils.Max(0, rightPath)

	return utils.Max(leftMax, rightMax, leftPath+rightPath+root.val),
		utils.Max(leftPath, rightPath) + root.val
}
