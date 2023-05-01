/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	// solution1: DFS，找到1之后找附近的1即可
	// solution2: 遍历，找到所有1
	counter := 0
	for rowIdx := 0; rowIdx < len(grid); rowIdx += 1 {
		for columnIdx := 0; columnIdx < len(grid[rowIdx]); columnIdx += 1 {
			if grid[rowIdx][columnIdx] == 0 {
				continue
			}

			// up
			if rowIdx == 0 || grid[rowIdx-1][columnIdx] == 0 {
				counter += 1
			}

			// right
			if columnIdx == len(grid[rowIdx])-1 || grid[rowIdx][columnIdx+1] == 0 {
				counter += 1
			}

			// bottom
			if rowIdx == len(grid)-1 || grid[rowIdx+1][columnIdx] == 0 {
				counter += 1
			}

			// left
			if columnIdx == 0 || grid[rowIdx][columnIdx-1] == 0 {
				counter += 1
			}
		}
	}
	return counter
}

// @lc code=end

