package id441

/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] Arranging Coins
 */

// @lc code=start
func arrangeCoins(n int) int {
	l := 1
	r := n
	for l <= r {
		m := l + (r-l)/2
		mm := (1 + m) * m / 2
		if mm < n {
			l = m + 1
		} else if mm > n {
			r = m - 1
		} else {
			return m
		}
	}
	return r
}

// @lc code=end
