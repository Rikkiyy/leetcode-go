package id349

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	var a [1000]int
	var b []int
	for _, v1 := range nums1 {
		a[v1] = 1
	}
	for _, v2 := range nums2 {
		if a[v2] == 1 {
			a[v2] = 2
			b = append(b, v2)
		}
	}
	return b
}

// @lc code=end
