/*
 * @lc app=leetcode.cn id=883 lang=golang
 *
 * [883] 三维形体投影面积
 */

// @lc code=start
func projectionArea(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	xy := 0
	xz := 0
	yz := 0

	for x := 0; x < len(grid); x += 1 {
		xyMaxVal := 0
		yxMaxVal := 0

		for y := 0; y < len(grid[0]); y += 1 {
			if grid[x][y] > 0 {
				xy += 1
			}

			if grid[x][y] > xyMaxVal {
				xyMaxVal = grid[x][y]
			}

			if grid[y][x] > yxMaxVal {
				yxMaxVal = grid[y][x]
			}
		}

		xz += xyMaxVal
		yz += yxMaxVal
	}

	return xy + xz + yz
}

// @lc code=end

