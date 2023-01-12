package offer

//面试题3：前n个数字二进制形式中1的个数
//题目：输入一个非负数n，请计算0到n之间每个数字的二进制形式中1的个数，并输出一个数组。
//例如，输入的n为4，由于0、1、2、3、4的二进制形式中1的个数分别为0、1、1、2、1，因此输出数组[0，1，1，2，1]。

// CountBits1 Using i&(i-1) to remove the rightest 1.
// Complexity: O(nk) where k refers to how many bits an int contains.
func CountBits1(N int) (res []int) {
	res = make([]int, N+1)
	for i := 0; i <= N; i++ {
		j := i
		for j > 0 {
			res[i]++
			j = j & (j - 1)
		}
	}
	return
}

// CountBits2 Using i&(i-1) to remove the rightest 1 too.
// However, it's kind of dynamic programming, where the number of ones of n would be that of n&(n-1)+1.
func CountBits2(N int) (res []int) {
	res = make([]int, N+1)
	for i := 1; i <= N; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return
}

// CountBits3 is still kind of dynamic programming.
func CountBits3(N int) (res []int) {
	res = make([]int, N+1)
	for i := 1; i <= N; i++ {
		res[i] = res[i>>1] + i&1
	}
	return
}
