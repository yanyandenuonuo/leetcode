/*
 * @lc app=leetcode.cn id=892 lang=golang
 *
 * [892] 三维形体的表面积
 */

// @lc code=start
func surfaceArea(grid [][]int) int {
	// solution: 计算每个点位表面积
	surface := 0

	for x := 0; x < len(grid); x += 1 {
		for y := 0; y < len(grid[0]); y += 1 {
			if grid[x][y] == 0 {
				continue
			}

			// top && bottom
			surface += 2

			// left
			if y == 0 {
				surface += grid[x][y]
			} else if grid[x][y] > grid[x][y-1] {
				surface += grid[x][y] - grid[x][y-1]
			}

			// right
			if y == len(grid[0])-1 {
				surface += grid[x][y]
			} else if grid[x][y] > grid[x][y+1] {
				surface += grid[x][y] - grid[x][y+1]
			}

			// backend
			if x == 0 {
				surface += grid[x][y]
			} else if grid[x][y] > grid[x-1][y] {
				surface += grid[x][y] - grid[x-1][y]
			}

			// front
			if x == len(grid)-1 {
				surface += grid[x][y]
			} else if grid[x][y] > grid[x+1][y] {
				surface += grid[x][y] - grid[x+1][y]
			}
		}
	}

	return surface
}

// @lc code=end

