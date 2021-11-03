package leetcode

// hashmap index
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	hashTable := make(map[int]int)
	// init map via shorter list
	for _, v := range nums2 {
		if _, ok := hashTable[v]; !ok {
			hashTable[v] = 1
		} else {
			hashTable[v]++
		}
	}
	var t []int
	for _, v := range nums1 {
		if numExist, ok := hashTable[v]; ok {
			if numExist > 0 {
				hashTable[v]--
				t = append(t, v)
			}
		}
	}
	return t
}
