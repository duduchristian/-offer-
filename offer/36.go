package offer

import "strconv"

//面试题36：后缀表达式
//题目：后缀表达式是一种算术表达式，它的操作符在操作数的后面。输入一个用字符串数组表示的后缀表达式，请输出该后缀表达式的计算结果。
//假设输入的一定是有效的后缀表达式。例如，后缀表达式["2"，"1"，"3"，"*"，"+"]对应的算术表达式是“2+1*3”，因此输出它的计算结果5。

func EvalRPN(tokens []string) int {
	stack := NewStack()
	for _, token := range tokens {
		switch token {
		case "+":
			fallthrough
		case "-":
			fallthrough
		case "*":
			fallthrough
		case "/":
			num1 := stack.Pop().(int)
			num2 := stack.Pop().(int)
			stack.Push(calculate(num1, num2, token))
		default:
			num, _ := strconv.ParseInt(token, 10, 64)
			stack.Push(int(num))
		}
	}
	return stack.Pop().(int)
}

func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 0
}
