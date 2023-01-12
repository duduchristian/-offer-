package offer

import (
	"unsafe"
)

//面试题4：只出现一次的数字
//题目：输入一个整数数组，数组中只有一个数字出现了一次，而其他数字都出现了3次。请找出那个只出现一次的数字。
//例如，如果输入的数组为[0，1，0，1，0，1，100]，则只出现一次的数字是100。

func SingleNumber(nums []int) (res int) {
	bits := int(unsafe.Sizeof(res) * 8)
	mem := make([]int, bits)
	for _, num := range nums {
		for i := 0; i < bits; i++ {
			mem[i] += (num >> i) & 1
		}
	}
	for i, count := range mem {
		if count%3 == 0 {
			continue
		}
		res += 1 << i
	}
	return
}

func SingleNumberBruteForce(nums []int) (res int) {
	mem := make(map[int]int, len(nums)/3+1)
	for _, num := range nums {
		mem[num] += 1
	}
	for k, v := range mem {
		if v != 3 {
			res = k
			break
		}
	}
	return
}
