/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	// solution: 前缀和+hash表, hash表的key为presum%k
	if len(nums) < 2 {
		return false
	}

	remainderMap := make(map[int]int, len(nums)+1)

	// 解决从0开始累加造成的不存在前缀和问题
	remainderMap[0] = -1

	for remainder, idx := 0, 0; idx < len(nums); idx += 1 {
		remainder = (remainder + nums[idx]) % k
		if preIdx, isExist := remainderMap[remainder]; isExist {
			if idx-preIdx >= 2 {
				return true
			}
		} else {
			// 注意一旦存在就不再更新idx以保证子数组更长
			remainderMap[remainder] = idx
		}
	}

	return false
}

// @lc code=end

