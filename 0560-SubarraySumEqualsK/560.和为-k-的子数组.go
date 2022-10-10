/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	// solution1: 逐次求解
	// res := 0
	// for idxStart := 0; idxStart < len(nums); idxStart += 1 {
	// 	subSum := 0
	// 	for idxEnd := idxStart; idxEnd < len(nums); idxEnd += 1 {
	// 		subSum += nums[idxEnd]
	// 		if subSum == k {
	// 			res += 1
	// 		}
	// 	}
	// }
	// return res

	// solution2: 前缀和
	res := 0
	preSum := 0
	sumMap := make(map[int]int, len(nums))

	// 当index = (0-1)时，认为累积和的值为0；如果不设置此值，也可在循环体内判断preSum = k时执行res+=1
	sumMap[0] = 1

	for idx := 0; idx < len(nums); idx += 1 {
		preSum += nums[idx]

		if _, isExist := sumMap[preSum-k]; isExist {
			// 存在index = x时， x+1到当前位置区间和为k
			res += sumMap[preSum-k]
		}
		sumMap[preSum] += 1
	}
	return res
}

// @lc code=end

