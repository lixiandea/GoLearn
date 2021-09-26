package main

func maxSubArray(nums []int) int {
	var sum int

	max := nums[0]

	index := 0
	// find first positive value
	for k, v := range nums {
		if v > max {
			max = v
		}
		if v >= 0 {
			max = v
			index = k
			break
		}
	}
	sum = max
	// fmt.Println(max)

	// iterate other num
	for _, v := range nums[index+1:] {
		sum += v
		if sum > 0 {
			if sum > max {
				max = sum
			}
		} else {
			sum = 0
		}
		// fmt.Println(max)
	}
	return max
}
