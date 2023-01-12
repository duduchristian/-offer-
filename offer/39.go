package offer

import "algorithm/utils"

//面试题39：直方图最大矩形面积

func LargestRectangleArea(heights []int) int {
	stack := NewStack()
	stack.Push(-1)

	maxArea := 0
	for i, height := range heights {
		for stack.PeekInt() != -1 && heights[stack.PeekInt()] >= height {
			h := heights[stack.PopInt()]
			w := i - stack.PeekInt() - 1
			maxArea = utils.Max(maxArea, h*w)
		}
		stack.PushInt(i)
	}

	for stack.PeekInt() != -1 {
		h := heights[stack.PopInt()]
		w := len(heights) - stack.PeekInt() - 1
		maxArea = utils.Max(maxArea, h*w)
	}

	return maxArea
}
