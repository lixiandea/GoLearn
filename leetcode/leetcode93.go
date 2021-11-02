package leetcode

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	rowContain, colContain := 1, 1
	for i := 0; i < len(matrix[0]); i++ {
		rowContain *= matrix[0][i]
	}
	for i := 0; i < len(matrix); i++ {
		colContain *= matrix[i][0]
	}
	// tag matrix
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			break
		}
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 {
				break
			}
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for k, v := range matrix[0] {
		if v == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][k] = 0
			}
		}
	}

	for j := 0; j < len(matrix); j++ {
		if matrix[j][0] == 0 {
			for i := 1; i < len(matrix[0]); i++ {
				matrix[j][i] = 0
			}
		}
	}

	if rowContain == 0 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if colContain == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
