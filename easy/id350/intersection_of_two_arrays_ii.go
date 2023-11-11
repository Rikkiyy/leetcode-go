package id350

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	var a [1001]int
	var b []int
	for _, v1 := range nums1 {
		a[v1] += 1
	}
	for _, v2 := range nums2 {
		if a[v2] > 0 {
			b = append(b, v2)
			a[v2] -= 1
		}
	}
	return b
}

// @lc code=end
