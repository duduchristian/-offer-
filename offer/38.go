package offer

//面试题38：每日温度
//题目：输入一个数组，它的每个数字是某天的温度。请计算每天需要等几天才会出现更高的温度
//例如，如果输入数组[35，31，33，36，34]，那么输出为[3，1，1，0，0]。
//由于第1天的温度是35℃，要等3天才会出现更高的温度36℃，因此对应的输出为3。
//第4天的温度是36℃，后面没有更高的温度，它对应的输出是0。其他的以此类推。

func DailyTemperatures(temperatures []int) []int {
	stack := NewStack()
	ret := make([]int, len(temperatures))
	for i, temp := range temperatures {
		for !stack.Empty() && temp > temperatures[stack.Peek().(int)] {
			index := stack.Pop().(int)
			ret[index] = i - index
		}
		stack.Push(i)
	}
	return ret
}
