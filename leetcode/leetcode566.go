package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	count := 0
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for count < r*c {
		res[count/c][count%c] = mat[count/n][count%n]
		count++
	}
	return res
}
