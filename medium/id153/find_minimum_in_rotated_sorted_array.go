package id153

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[l] <= nums[r] {
			return nums[l]
		} else {
			if nums[m] > nums[l] {
				l = m + 1
			} else if nums[m] < nums[l] {
				r = m
			} else {
				return nums[r]
			}
		}
	}
	return nums[r]
}

// @lc code=end
