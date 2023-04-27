/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	// solution: 利用位排序思想，因为num范围不定，所以需要通过map的方式来做
	if len(nums) == 0 {
		return 0
	}

	numMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		numMap[num] = true
	}

	maxLength := 1

	for num := range numMap {
		if numMap[num-1] || !numMap[num+1] {
			// 不是连续序列的最小值
			continue
		}

		nextNum := num + 1
		for numMap[nextNum] {
			nextNum += 1
		}

		if nextNum-num > maxLength {
			maxLength = nextNum - num
		}

		if maxLength > (len(numMap)+1)/2 {
			break
		}
	}

	return maxLength
}

// @lc code=end

