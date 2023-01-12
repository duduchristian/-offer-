package offer

import (
	"algorithm/utils"
	"strings"
)

//面试题2：二进制加法
//题目：输入两个表示二进制的字符串，请计算它们的和，并以二进制字符串的形式输出。例如，输入的二进制字符串分别是"11"和"10"，则输出"101"。

func AddBinary(a, b string) string {
	builder := strings.Builder{}
	i, j := len(a)-1, len(b)-1
	var carry uint8 = 0
	for i >= 0 || j >= 0 {
		res := carry
		if i >= 0 {
			res += a[i] - '0'
		}
		if j >= 0 {
			res += b[j] - '0'
		}
		carry = res / 2
		res %= 2
		builder.WriteByte('0' + res)
		i--
		j--
	}

	if carry == 1 {
		builder.WriteByte('1')
	}
	s := builder.String()
	utils.RevertStringBytes(s)
	return s
}
