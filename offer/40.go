package offer

import "algorithm/utils"

//面试题40：矩阵中的最大矩形

func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0]))
	maxArea := 0
	for _, row := range matrix {
		for i, num := range row {
			if num == '0' {
				heights[i] = 0
			} else {
				heights[i]++
			}
		}
		maxArea = utils.Max(maxArea, LargestRectangleArea(heights))
	}

	return maxArea
}
