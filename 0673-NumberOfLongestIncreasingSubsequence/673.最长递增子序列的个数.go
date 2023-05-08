/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	// solution1: 动态规划
	maxLIS, maxFreq := 0, 0
	dp, freqCounter := make([]int, len(nums)), make([]int, len(nums))

	for idx := 0; idx < len(nums); idx += 1 {
		dp[idx], freqCounter[idx] = 1, 1
		for leftIdx := 0; leftIdx < idx; leftIdx += 1 {
			if nums[idx] > nums[leftIdx] {
				if dp[leftIdx]+1 > dp[idx] {
					dp[idx], freqCounter[idx] = dp[leftIdx]+1, freqCounter[leftIdx]
				} else if dp[leftIdx]+1 == dp[idx] {
					freqCounter[idx] += freqCounter[leftIdx]
				}
			}
		}

		if dp[idx] > maxLIS {
			maxLIS = dp[idx]
			maxFreq = freqCounter[idx]
		} else if dp[idx] == maxLIS {
			maxFreq += freqCounter[idx]
		}
	}

	return maxFreq

	// solution2: 贪心+前缀和+二分查找
	// todo
}

// @lc code=end

