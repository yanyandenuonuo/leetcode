/*
 * @lc app=leetcode.cn id=910 lang=golang
 *
 * [910] 最小差值 II
 */

// @lc code=start
func smallestRangeII(nums []int, k int) int {
	// solution: 排序，排序后较小值+k，较大值-k
	sort.Ints(nums)
	minDiff := nums[len(nums)-1] - nums[0]
	for idx := 0; idx < len(nums)-1; idx += 1 {
		// 最大值为max(当前值+k, 最大值-k)
		maxVal := nums[len(nums)-1] - k
		if nums[idx]+k > maxVal {
			maxVal = nums[idx] + k
		}

		// 最小值为min(下一个值-k, 最小值+k)
		minVal := nums[idx+1] - k
		if nums[0]+k < minVal {
			minVal = nums[0] + k
		}

		if maxVal-minVal < minDiff {
			minDiff = maxVal - minVal
		}
	}

	return minDiff
}

// @lc code=end

