/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for idxX := range grid {
		for idxY := range grid[idxX] {
			if grid[idxX][idxY] == 0 {
				continue
			}
			area := areaIsLand(grid, idxX, idxY)

			if maxArea < area {
				maxArea = area
			}
		}
	}
	return maxArea
}

func areaIsLand(grid [][]int, row, column int) int {
	area := 0
	if grid[row][column] == 0 {
		return area
	}
	grid[row][column] = 0
	area += 1

	// up
	if row-1 >= 0 && grid[row-1][column] == 1 {
		area += areaIsLand(grid, row-1, column)
	}

	// down
	if row+1 < len(grid) && grid[row+1][column] == 1 {
		area += areaIsLand(grid, row+1, column)
	}

	// left
	if column-1 >= 0 && grid[row][column-1] == 1 {
		area += areaIsLand(grid, row, column-1)
	}

	// right
	if column+1 < len(grid[row]) && grid[row][column+1] == 1 {
		area += areaIsLand(grid, row, column+1)
	}
	return area
}

// @lc code=end

