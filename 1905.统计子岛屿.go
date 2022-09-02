/*
 * @lc app=leetcode.cn id=1905 lang=golang
 *
 * [1905] 统计子岛屿
 */

// @lc code=start
func countSubIslands(grid1 [][]int, grid2 [][]int) int {

	// [
	// 	[1,1,1,0,0],
	// 	[0,1,1,1,1],
	// 	[0,0,0,0,0],
	// 	[1,0,0,0,0],
	// 	[1,1,0,1,1]
	// ]

	// [
	// 	[1,1,1,0,0],
	// 	[0,0,1,1,1],
	// 	[0,1,0,0,0],
	// 	[1,0,1,1,0],
	// 	[0,1,0,1,0]
	// ]

	totalSubIsLand := 0
	for row := range grid2 {
		for column := range grid2[row] {
			if grid2[row][column] == 0 {
				continue
			}
			if isSubIsLand(grid2, grid1, row, column) {

				totalSubIsLand += 1
			}
		}
	}
	return totalSubIsLand
}

func isSubIsLand(grid, grid1 [][]int, row, column int) bool {
	if grid[row][column] == 0 {
		return false
	}

	grid[row][column] = 0

	isSub := true

	if grid1[row][column] == 0 {
		isSub = false
	}

	// up
	if row-1 >= 0 && grid[row-1][column] == 1 {
		isSub = isSubIsLand(grid, grid1, row-1, column) && isSub
	}

	// down
	if row+1 < len(grid) && grid[row+1][column] == 1 {
		isSub = isSubIsLand(grid, grid1, row+1, column) && isSub
	}

	// left
	if column-1 >= 0 && grid[row][column-1] == 1 {
		isSub = isSubIsLand(grid, grid1, row, column-1) && isSub
	}

	// right
	if column+1 < len(grid[row]) && grid[row][column+1] == 1 {
		isSub = isSubIsLand(grid, grid1, row, column+1) && isSub
	}

	return isSub
}

// @lc code=end

