/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	// solution1: 动态规划
	/**
	dp := make([]int, len(nums))
	for idx := 0; idx < len(nums); idx += 1 {
		for step := 1; step <= nums[idx] && idx+step < len(nums); step += 1 {
			if dp[idx+step] == 0 || dp[idx+step] > 1+dp[idx] {
				dp[idx+step] = 1 + dp[idx]
			}
		}
	}

	return dp[len(nums)-1]
	*/

	// solution2: 贪心，局部最优达到全局最优
	counter := 0       // 计数
	currentMaxIdx := 0 // 当前一步能达到的最大Idx
	preMaxIdx := 0     // 前一步能达到的最大Idx

	for idx := 0; idx < len(nums)-1; idx += 1 {
		if idx+nums[idx] > currentMaxIdx {
			currentMaxIdx = idx + nums[idx]
		}

		if idx == preMaxIdx {
			preMaxIdx = currentMaxIdx
			counter += 1
		}
	}

	return counter
}

// @lc code=end

