/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	// solution: 动态规划
	dp := make([][]int, len(grid))
	for rowIdx := 0; rowIdx < len(grid); rowIdx += 1 {
		dp[rowIdx] = make([]int, len(grid[rowIdx]))
		for columnIdx := 0; columnIdx < len(grid[rowIdx]); columnIdx += 1 {
			if rowIdx == 0 && columnIdx == 0 {
				dp[rowIdx][columnIdx] = grid[rowIdx][columnIdx]
			} else if rowIdx == 0 && columnIdx != 0 {
				dp[rowIdx][columnIdx] = dp[rowIdx][columnIdx-1] + grid[rowIdx][columnIdx]
			} else if rowIdx != 0 && columnIdx == 0 {
				dp[rowIdx][columnIdx] = dp[rowIdx-1][columnIdx] + grid[rowIdx][columnIdx]
			} else {
				if dp[rowIdx-1][columnIdx] < dp[rowIdx][columnIdx-1] {
					dp[rowIdx][columnIdx] = dp[rowIdx-1][columnIdx] + grid[rowIdx][columnIdx]
				} else {
					dp[rowIdx][columnIdx] = dp[rowIdx][columnIdx-1] + grid[rowIdx][columnIdx]
				}
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

// @lc code=end

