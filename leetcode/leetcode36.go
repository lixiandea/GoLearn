package leetcode

func isValidSudoku(board [][]byte) bool {
	var ver [9][9]int
	var her [9][9]int
	var box [9][9]int
	for k1, v1 := range board {
		for k2, v2 := range v1 {
			if v2 == '.' {
				continue
			}
			cur := int(v2 - '1')
			if ver[k1][cur] != 0 {
				return false
			}
			if her[k2][cur] != 0 {
				return false
			}
			if box[k1/3+k2/3*3][cur] != 0 {
				return false
			}
			ver[k1][cur]++
			her[k2][cur]++
			box[k1/3+(k2/3)*3][cur]++
		}
	}
	return true
}
