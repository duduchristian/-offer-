package offer

//试题13：二维子矩阵的数字之和
//题目：输入一个二维矩阵，如何计算给定左上角坐标和右下角坐标的子矩阵的数字之和？对于同一个二维矩阵，
//计算子矩阵的数字之和的函数可能由于输入不同的坐标而被反复调用多次。例如，输入图2.1中的二维矩阵，
//以及左上角坐标为（2，1）和右下角坐标为（4，3）的子矩阵，该函数输出8。

type SumMatrix [][]int

func NumMatrix(matrix SumMatrix) SumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	M, N := len(matrix), len(matrix[0])

	sums := make([][]int, M+1)
	for i := range sums {
		sums[i] = make([]int, N+1)
	}

	for i := 0; i < M; i++ {
		rowSum := 0
		for j := 0; j < N; j++ {
			rowSum += matrix[i][j]
			sums[i+1][j+1] = sums[i][j+1] + rowSum
		}
	}
	return sums
}

func (m SumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return m[row2+1][col2+1] - m[row1][col2+1] - m[row2+1][col1] + m[row1][col1]
}
