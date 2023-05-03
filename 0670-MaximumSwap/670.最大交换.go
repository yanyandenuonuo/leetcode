/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start
func maximumSwap(num int) int {
	// solution1: 枚举
	// solution2: 贪心：左边数字越小越靠左，右边数字越大越靠右
	numByte := []byte(strconv.Itoa(num))
	leftIdx, rightIdx, maxIdx := -1, -1, len(numByte)-1
	for idx := len(numByte) - 2; idx >= 0; idx -= 1 {
		if numByte[idx] > numByte[maxIdx] {
			maxIdx = idx
			continue
		}
		if numByte[idx] < numByte[maxIdx] {
			leftIdx, rightIdx = idx, maxIdx
		}
	}

	if leftIdx != rightIdx {
		numByte[leftIdx], numByte[rightIdx] = numByte[rightIdx], numByte[leftIdx]
	}

	newNum, _ := strconv.Atoi(string(numByte))

	return newNum
}

// @lc code=end

