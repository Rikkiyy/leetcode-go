package id69

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	var r int = x
	l := 1
	for l <= r {
		var m int = l + (r-l)/2
		if x/m > m {
			l = m + 1
		} else if x/m < m {
			r = m - 1
		} else {
			return m
		}
	}
	return r
}

// @lc code=end
