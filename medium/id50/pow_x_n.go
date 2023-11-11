package id50

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	isN := false
	if n < 0 {
		isN = true
		n = -n
	}
	r := strconv.FormatInt(int64(n), 2)
	ans := float64(1)
	x_temp := x
	for i := len(r) - 1; i >= 0; i-- {
		if int(r[i]) == 49 {
			ans = ans * x_temp
		}
		if i == 0 {
			break
		}
		x_temp *= x_temp
	}
	if isN {
		ans = 1 / ans
	}
	return ans
}

// @lc code=end
