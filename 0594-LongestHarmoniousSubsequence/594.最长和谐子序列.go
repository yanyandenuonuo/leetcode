/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

// @lc code=start
func findLHS(nums []int) int {
	// solution1: 排序+滑动窗口
	// solution2: hash
	numMap := make(map[int]int, 0)
	for _, num := range nums {
		numMap[num] += 1
	}

	maxLHS := 0
	for num := range numMap {
		if _, isExist := numMap[num+1]; isExist && numMap[num]+numMap[num+1] > maxLHS {
			maxLHS = numMap[num] + numMap[num+1]
		}
	}

	return maxLHS
}

// @lc code=end

