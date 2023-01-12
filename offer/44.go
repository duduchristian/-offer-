package offer

import (
	"algorithm/utils"
	"math"
)

func LargestValues(root TreeNode) []int {
	queue1 := NewQueue()
	queue2 := NewQueue()
	if root != nil {
		queue1.Offer(root)
	}

	var result []int
	max := math.MinInt
	for !queue1.Empty() {
		node := queue1.Poll().(TreeNode)
		max = utils.Max(max, node.val)
		if node.left != nil {
			queue2.Offer(node.left)
		}
		if node.right != nil {
			queue2.Offer(node.right)
		}
		if queue1.Empty() {
			result = append(result, max)
			max = math.MinInt

			queue1, queue2 = queue2, queue1
		}
	}
	return result
}
