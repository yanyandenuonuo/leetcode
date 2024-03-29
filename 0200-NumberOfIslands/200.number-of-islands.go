/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (47.41%)
 * Likes:    6937
 * Dislikes: 221
 * Total Accepted:    871K
 * Total Submissions: 1.8M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given an m x n 2d grid map of '1's (land) and '0's (water), return the
 * number of islands.
 * 
 * An island is surrounded by water and is formed by connecting adjacent lands
 * horizontally or vertically. You may assume all four edges of the grid are
 * all surrounded by water.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * Output: 3
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 300
 * grid[i][j] is '0' or '1'.
 * 
 * 
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	res := 0
	for row := 0; row < len(grid); row += 1 {
		for column := 0; column < len(grid[row]); column += 1 {
			if grid[row][column] == '1' {
				res += 1
				dfs(grid, row, column)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, row int, column int) {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) {
		return
	}
	if grid[row][column] != '1' {
		return
	}
	grid[row][column] = '2'

	// up
	dfs(grid, row-1, column)

	// down
	dfs(grid, row+1, column)

	// left
	dfs(grid, row, column-1)

	// right
	dfs(grid, row, column+1)
}

// @lc code=end

