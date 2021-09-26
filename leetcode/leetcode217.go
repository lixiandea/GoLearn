// package leetcode

// // one way solve leetcode 217
// func containsDuplicate(nums []int) bool {
// 	set := make(map[int]int, 1024)

// 	for _, k := range nums {
// 		if _, ok := set[k]; ok {
// 			return true
// 		} else {
// 			set[k] = 1
// 		}
// 	}
// 	return false
// }
