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

	memeryGrid := make([][]int, len(grid))

	freshOrangeCount := 0

	for row := range grid {
		memeryGrid[row] = make([]int, len(grid[row]))
		for column := range grid[row] {
			switch grid[row][column] {
			case 0:
				continue
			case 1:
				freshOrangeCount += 1
				memeryGrid[row][column] = -1
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

	currentQueueCount := len(queueElement)
	nextQueueCount := 0

	for len(queueElement) > 0 {
		row := queueElement[0][0]
		column := queueElement[0][1]

		queueElement = queueElement[1:]

		// up
		if row-1 >= 0 && memeryGrid[row-1][column] == -1 {
			freshOrangeCount -= 1

			memeryGrid[row-1][column] = 0
			queueElement = append(queueElement, []int{row - 1, column})
			nextQueueCount += 1
		}

		// down
		if row+1 < len(grid) && memeryGrid[row+1][column] == -1 {
			freshOrangeCount -= 1

			memeryGrid[row+1][column] = 0
			queueElement = append(queueElement, []int{row + 1, column})
			nextQueueCount += 1
		}

		// left
		if column-1 >= 0 && memeryGrid[row][column-1] == -1 {
			freshOrangeCount -= 1

			memeryGrid[row][column-1] = 0
			queueElement = append(queueElement, []int{row, column - 1})
			nextQueueCount += 1
		}

		// right
		if column+1 < len(grid[row]) && memeryGrid[row][column+1] == -1 {
			freshOrangeCount -= 1

			memeryGrid[row][column+1] = 0
			queueElement = append(queueElement, []int{row, column + 1})
			nextQueueCount += 1
		}

		currentQueueCount -= 1

		if currentQueueCount == 0 {
			timeCount += 1
			currentQueueCount = nextQueueCount
			nextQueueCount = 0
		}
	}

	if freshOrangeCount > 0 {
		return -1
	}

	return timeCount - 1
}

// @lc code=end

