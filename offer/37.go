package offer

//面试题37：小行星碰撞

func AsteroidCollision(asteroids []int) []int {
	stack := NewStack()
	for _, as := range asteroids {
		for !stack.Empty() && stack.Peek().(int) > 0 && stack.Peek().(int) < -as {
			stack.Pop()
		}

		if !stack.Empty() && as < 0 && stack.Peek().(int) == -as {
			stack.Pop()
		} else if as > 0 || stack.Empty() || stack.Peek().(int) < 0 {
			stack.Push(as)
		}
	}
	ret := make([]int, stack.Len())
	for i := range ret {
		ret[i] = stack.Pop().(int)
	}
	return ret
}
