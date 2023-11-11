package id278

/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l := 1
	r := n
	for l <= r {
		m := l + (r-l)/2
		if isBadVersion(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

// @lc code=end
