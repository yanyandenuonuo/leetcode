/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 *
 * https://leetcode.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (37.99%)
 * Likes:    3581
 * Dislikes: 93
 * Total Accepted:    293.8K
 * Total Submissions: 773K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * Given a 2D binary matrix filled with 0's and 1's, find the largest square
 * containing only 1's and return its area.
 *
 * Example:
 *
 *
 * Input:
 *
 * 1 0 1 0 0
 * 1 0 1 1 1
 * 1 1 1 1 1
 * 1 0 0 1 0
 *
 * Output: 4
 *
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	// solution1: 暴力穷举边长，超时
	/**
	maxSideLength := 1
	for rowIdx := 0; rowIdx < len(matrix); rowIdx += 1 {
		for columnIdx := 0; columnIdx < len(matrix[rowIdx]); columnIdx += 1 {
			if matrix[rowIdx][columnIdx] == '0' {
				continue
			}
			for sideLength := 2; sideLength <= len(matrix); sideLength += 1 {
				if !isSquare(matrix, rowIdx, columnIdx, sideLength) {
					break
				}
				if sideLength > maxSideLength {
					maxSideLength = sideLength
				}
			}
		}
	}
	return maxSideLength * maxSideLength
	*/

	// solution2: 动态规划
	//			  matrix[rowIdx][columnIdx] == 0, dp[rowIdx][columnIdx] = 0
	//			  matrix[rowIdx][columnIdx] == 1, dp[rowIdx][columnIdx] = min(dp[rowIdx-1][columnIdx-1], dp[rowIdx-1][columnIdx], dp[rowIdx][columnIdx-1])
	dp := make([][]int, len(matrix))
	for rowIdx := 0; rowIdx < len(matrix); rowIdx += 1 {
		dp[rowIdx] = make([]int, len(matrix[rowIdx]))
	}
	maxSideLength := 0

	for rowIdx := 0; rowIdx < len(matrix); rowIdx += 1 {
		for columnIdx := 0; columnIdx < len(matrix[rowIdx]); columnIdx += 1 {
			if matrix[rowIdx][columnIdx] == '0' {
				dp[rowIdx][columnIdx] = 0
				continue
			}
			if rowIdx == 0 || columnIdx == 0 {
				dp[rowIdx][columnIdx] = 1
			} else {
				dp[rowIdx][columnIdx] = len(matrix)
				if dp[rowIdx-1][columnIdx-1] < dp[rowIdx][columnIdx] {
					dp[rowIdx][columnIdx] = dp[rowIdx-1][columnIdx-1]
				}

				if dp[rowIdx-1][columnIdx] < dp[rowIdx][columnIdx] {
					dp[rowIdx][columnIdx] = dp[rowIdx-1][columnIdx]
				}

				if dp[rowIdx][columnIdx-1] < dp[rowIdx][columnIdx] {
					dp[rowIdx][columnIdx] = dp[rowIdx][columnIdx-1]
				}

				dp[rowIdx][columnIdx] += 1
			}

			if dp[rowIdx][columnIdx] > maxSideLength {
				maxSideLength = dp[rowIdx][columnIdx]
			}
		}
	}
	return maxSideLength * maxSideLength
}

func isSquare(matrix [][]byte, rowStartIdx, columnStartIdx, sideLength int) bool {
	if rowStartIdx+sideLength-1 >= len(matrix) {
		return false
	}

	if columnStartIdx+sideLength-1 >= len(matrix[rowStartIdx]) {
		return false
	}

	for rowIdx := rowStartIdx; rowIdx < rowStartIdx+sideLength; rowIdx += 1 {
		if matrix[rowIdx][columnStartIdx+sideLength-1] == '0' {
			return false
		}
	}

	for columnIdx := columnStartIdx; columnIdx < columnStartIdx+sideLength; columnIdx += 1 {
		if matrix[rowStartIdx+sideLength-1][columnIdx] == '0' {
			return false
		}
	}

	return true
}

// @lc code=end

