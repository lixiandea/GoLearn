package leetcode

func twoSum(nums []int, target int) []int {
	t := make(map[int]int)
	for k, v := range nums {
		if index, isContain := t[target-v]; isContain {
			return []int{k, index}
		} else {
			t[v] = k
		}
	}
	return nil
}
