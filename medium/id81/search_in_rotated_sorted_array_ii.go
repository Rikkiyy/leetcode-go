package id81

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[l] < nums[r] {
			if target > nums[m] {
				l = m + 1
			} else if target < nums[m] {
				r = m - 1
			} else {
				return true
			}
		} else if nums[l] == nums[r] {
			if target == nums[m] || target == nums[l] {
				return true
			} else {
				l = l + 1
				r = r - 1
			}
		} else {
			if nums[m] < nums[r] {
				if target > nums[m] && target <= nums[r] {
					l = m + 1
				} else if target == nums[m] {
					return true
				} else {
					r = m - 1
				}
			} else if nums[m] == nums[r] {
				if target == nums[m] {
					return true
				} else {
					r = m - 1
				}
			} else {
				if target < nums[m] && target >= nums[l] {
					r = m - 1
				} else if target == nums[m] {
					return true
				} else {
					l = m + 1
				}
			}
		}
	}
	return false
}

// @lc code=end
