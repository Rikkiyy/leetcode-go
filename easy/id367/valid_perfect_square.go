package id367

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	a := 1
	b := num
	for a <= b {
		m := a + (b-a)/2
		if num/m > m {
			a = m + 1
		} else if num/m < m {
			b = m - 1
		} else if num%m == 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

// @lc code=end
