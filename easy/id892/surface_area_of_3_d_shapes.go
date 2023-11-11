/*
 * @lc app=leetcode.cn id=892 lang=golang
 *
 * [892] Surface Area of 3D Shapes
 */

// @lc code=start
func surfaceArea(grid [][]int) int {
	a := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if j == (len(grid) - 1) {
				a += grid[i][j]
			} else if (grid[i][j] - grid[i][j+1]) > 0 {
				a += (grid[i][j] - grid[i][j+1])
			}
			if i == (len(grid) - 1) {
				a += grid[i][j]
			} else if (grid[i][j] - grid[i+1][j]) > 0 {
				a += (grid[i][j] - grid[i+1][j])
			}
			if grid[i][j] > 0 {
				a += 1
			}
		}
	}
	return a * 2
}

// @lc code=end

