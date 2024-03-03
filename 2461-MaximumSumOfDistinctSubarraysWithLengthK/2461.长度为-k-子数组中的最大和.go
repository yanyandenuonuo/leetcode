/*
 * @lc app=leetcode.cn id=2461 lang=golang
 *
 * [2461] 长度为 K 子数组中的最大和
 */

// @lc code=start
func maximumSubarraySum(nums []int, k int) int64 {
	sumNums := make([]int64, len(nums))
	sumNums[0] = int64(nums[0])
	for idx := 1; idx < len(nums); idx += 1 {
		sumNums[idx] = sumNums[idx-1] + int64(nums[idx])
	}

	subNumMap := make(map[int]int, k)
	for leftIdx := 0; leftIdx < k; leftIdx += 1 {
		subNumMap[nums[leftIdx]] += 1
	}

	maxSubSum := int64(0)
	for leftIdx := 0; leftIdx <= len(nums)-k; leftIdx += 1 {
		if len(subNumMap) == k {
			preSum := int64(0)

			if leftIdx > 0 {
				preSum = sumNums[leftIdx-1]
			}

			subSum := sumNums[leftIdx+k-1] - preSum
			if subSum > maxSubSum {
				maxSubSum = subSum
			}
		}

		subNumMap[nums[leftIdx]] -= 1
		if subNumMap[nums[leftIdx]] == 0 {
			delete(subNumMap, nums[leftIdx])
		}

		if leftIdx+k < len(nums) {
			subNumMap[nums[leftIdx+k]] += 1
		}
	}

	return maxSubSum
}

// @lc code=end

