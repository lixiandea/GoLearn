package leetcode

// double Pointer
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	for n > 0 {
		if m == 0 {
			nums1[i] = nums2[n-1]
			n--
		} else {
			if nums2[n-1] > nums1[m-1] {
				nums1[i] = nums2[n-1]
				n--
			} else {
				nums1[i] = nums1[m-1]
				m--
			}
		}
		i--
	}
}
