/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	maxRow := len(triangle) - 1
	minNum := 0
	memoryMap := make(map[int]map[int]int, len(triangle)*len(triangle[maxRow]))
	for idx := 0; idx < len(triangle[maxRow]); idx += 1 {
		currentNum := dp(triangle, memoryMap, maxRow, idx)
		if idx == 0 || currentNum < minNum {
			minNum = currentNum
		}
	}
	return minNum
}

func dp(triangle [][]int, memoryMap map[int]map[int]int, row, column int) int {
	if _, isSet := memoryMap[row][column]; isSet {
		return memoryMap[row][column]
	}

	if row == 0 {
		return triangle[row][column]
	}
	left := 0
	existLeft := false
	existRight := false
	if column-1 >= 0 {
		existLeft = true
		left = triangle[row][column] + dp(triangle, memoryMap, row-1, column-1)
	}

	right := -1
	if column < len(triangle[row-1]) {
		existRight = true
		right = triangle[row][column] + dp(triangle, memoryMap, row-1, column)
	}
	if _, isSet := memoryMap[row]; !isSet {
		memoryMap[row] = make(map[int]int, len(triangle[row]))
	}
	if existLeft && existRight {
		if left < right {
			memoryMap[row][column] = left
			return left
		}
		memoryMap[row][column] = right
		return right
	} else if existLeft {
		memoryMap[row][column] = left
		return left
	} else if existRight {
		memoryMap[row][column] = right
		return right
	}

	// exception
	return -1
}

// @lc code=end

