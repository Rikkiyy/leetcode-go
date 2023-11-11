package id392

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)
	if len(sb) == 0 {
		return true
	}
	for i, j := 0, 0; i <= len(sb) && j < len(tb); j++ {
		if tb[j] == sb[i] {
			i++
		}
		if i == len(sb) {
			return true
		}
	}
	return false

}

// @lc code=end
