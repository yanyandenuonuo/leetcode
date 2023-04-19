/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	maxVal := 0
	for _, val := range nums {
		sum += val
		if val > maxVal {
			maxVal = val
		}
	}

	// sum = partition1 + partition2 && partition1 == partition2
	if sum%2 != 0 {
		return false
	}

	targetSum := sum / 2
	if maxVal > targetSum {
		return false
	}

	/**
	dp := make([][]bool, len(nums))

	for idx, val := range nums {
		dp[idx] = make([]bool, targetSum+1)
		dp[idx][val] = true
	}

	for idx := 1; idx < len(nums); idx += 1 {
		for subSum := 0; subSum <= targetSum; subSum += 1 {
			if nums[idx] < subSum {
				dp[idx][subSum] = dp[idx-1][subSum] || dp[idx-1][subSum-nums[idx]]
			} else if nums[idx] > subSum {
				dp[idx][subSum] = dp[idx-1][subSum]
			} else {
				dp[idx][subSum] = true
			}
		}

		if dp[len(nums)-1][targetSum] {
			return true
		}
	}

	return dp[len(nums)-1][targetSum]
	*/

	dp := make([]bool, targetSum+1)

	dp[0] = true

	// dp[idx][sum] = dp[idx-1][sum] || dp[idx-1][sum-num[idx]]
	// 因为dp[idx]只使用了dp[idx-1]的值，则可以做简化处理
	// 倒序赋值，则可以保证dp[sum-num[idx]]的值实际为dp[idx-1][sum-num[idx]]
	for _, val := range nums {
		for subSum := targetSum; subSum >= val; subSum -= 1 {
			dp[subSum] = dp[subSum] || dp[subSum-val]
		}

		if dp[targetSum] {
			return true
		}
	}

	return dp[targetSum]
}

// @lc code=end

