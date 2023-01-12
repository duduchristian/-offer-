package offer

import "math"

//面试题1：整数除法
//题目：输入2个int型整数，它们进行除法计算并返回商，要求不得使用乘号'*'、除号'/'及求余符号'%'。
//当发生溢出时，返回最大的整数值。假设除数不为0。例如，输入15和2，输出15/2的结果，即7。

func Divide(dividend, divisor int) int {
	if dividend == math.MinInt && divisor == -1 {
		return math.MaxInt
	}

	negative := 2
	if dividend > 0 {
		negative--
		dividend = -dividend
	}

	if divisor > 0 {
		negative--
		divisor = -divisor
	}

	if negative == 1 {
		return -divideCore(dividend, divisor)
	}
	return divideCore(dividend, divisor)
}

func divideCore(dividend, divisor int) (result int) {
	for dividend <= divisor {
		value := divisor
		quotient := 1
		for value >= -0x4000000000000000 && dividend <= value+value {
			quotient += quotient
			value += value
		}
		result += quotient
		dividend -= value
	}
	return
}
