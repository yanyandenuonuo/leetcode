/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	// solution: hash
	numIdxMap, numCounterMap := make(map[int]int), make(map[int]int)
	maxFreq, minLength := 0, len(nums)
	for idx, num := range nums {
		if _, isExist := numIdxMap[num]; !isExist {
			numIdxMap[num] = idx
		}

		numCounterMap[num] += 1

		if numCounterMap[num] > maxFreq {
			maxFreq = numCounterMap[num]
			minLength = idx - numIdxMap[num] + 1
		} else if numCounterMap[num] == maxFreq {
			if idx-numIdxMap[num]+1 < minLength {
				minLength = idx - numIdxMap[num] + 1
			}
		}
	}

	return minLength
}

// @lc code=end

