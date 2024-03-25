/*
 * @lc app=leetcode.cn id=1375 lang=golang
 *
 * [1375] 二进制字符串前缀一致的次数
 */

// @lc code=start
func numTimesAllBlue(flips []int) int {
	// solution1: 前缀和等于index自然增长之和
	/**
	idxSum, flipSum, counter := 0, 0, 0
	for idx, flip := range flips {
		idxSum += idx + 1
		flipSum += flip

		if idxSum == flipSum {
			counter += 1
		}
	}

	return counter
	*/

	// solution2: 找最大值，当最大值等于index+1时，说明前缀一致，和思路一较为相似
	maxflip, counter := 0, 0
	for idx, flip := range flips {
		if flip > maxflip {
			maxflip = flip
		}

		if maxflip == idx+1 {
			counter += 1
		}
	}

	return counter
}

// @lc code=end

