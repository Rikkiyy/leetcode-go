package id33

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[l] <= nums[r] {
			if target > nums[m] {
				l = m + 1
			} else if target < nums[m] {
				r = m - 1
			} else {
				return m
			}
		} else {
			if nums[m] < nums[r] {
				if target > nums[m] && target <= nums[r] {
					l = m + 1
				} else if target == nums[m] {
					return m
				} else {
					r = m - 1
				}
			} else {
				if target < nums[m] && target >= nums[l] {
					r = m - 1
				} else if target == nums[m] {
					return m
				} else {
					l = m + 1
				}
			}
		}
	}
	return -1
}

// @lc code=end
