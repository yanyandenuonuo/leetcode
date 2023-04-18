/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start

func orangesRotting(grid [][]int) int {
	// 0 empty
	// 1 fresh
	// 2 rotten

	queueElement := make([][]int, 0, len(grid)*len(grid[0])-1)

	freshOrangeCount := 0

	for row := range grid {
		for column := range grid[row] {
			switch grid[row][column] {
			case 0:
				continue
			case 1:
				freshOrangeCount += 1
			case 2:
				queueElement = append(queueElement, []int{row, column})
			default:
				continue
			}
		}
	}

	if freshOrangeCount == 0 {
		return 0
	}

	timeCount := 0

	nextQueue := make([][]int, 0, len(queueElement))

	for len(queueElement) > 0 {
		row := queueElement[0][0]
		column := queueElement[0][1]

		// up
		if row-1 >= 0 && grid[row-1][column] == 1 {
			freshOrangeCount -= 1

			grid[row-1][column] = 2
			nextQueue = append(nextQueue, []int{row - 1, column})
		}

		// down
		if row+1 < len(grid) && grid[row+1][column] == 1 {
			freshOrangeCount -= 1

			grid[row+1][column] = 2
			nextQueue = append(nextQueue, []int{row + 1, column})
		}

		// left
		if column-1 >= 0 && grid[row][column-1] == 1 {
			freshOrangeCount -= 1

			grid[row][column-1] = 2
			nextQueue = append(nextQueue, []int{row, column - 1})
		}

		// right
		if column+1 < len(grid[row]) && grid[row][column+1] == 1 {
			freshOrangeCount -= 1

			grid[row][column+1] = 2
			nextQueue = append(nextQueue, []int{row, column + 1})
		}

		if len(queueElement) == 1 {
			timeCount += 1
			queueElement = nextQueue
			nextQueue = make([][]int, 0, len(queueElement))
		} else {
			queueElement = queueElement[1:]
		}
	}

	if freshOrangeCount > 0 {
		return -1
	}

	return timeCount - 1
}

// @lc code=end

