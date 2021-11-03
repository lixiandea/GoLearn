package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := 10000
	res := 0
	for _, v := range prices {
		if v < min {
			min = v
		}
		if res < v-min {
			res = v - min
		}
	}
	return res
}
