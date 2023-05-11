/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	// solution: 滑动窗口
	maxSum, sum := -1<<31, 0
	for idx := 0; idx < len(nums); idx += 1 {
		if idx < k {
			sum += nums[idx]
			continue
		} else if idx == k {
			if sum > maxSum {
				maxSum = sum
			}
		}

		sum -= nums[idx-k]
		sum += nums[idx]
		if sum > maxSum {
			maxSum = sum
		}
	}

	if len(nums) == k {
		maxSum = sum
	}

	return float64(maxSum) / float64(k)
}

// @lc code=end

