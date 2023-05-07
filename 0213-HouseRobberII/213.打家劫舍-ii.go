/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	// solution: 动态规划，因为房屋首尾相连，则可以选择nums[:len(nums)-1]和nums[1:]的较大值
	if len(nums) == 1 {
		return nums[0]
	}

	stolenFirst := robHouse(nums, 0, len(nums)-2)
	stolenEnd := robHouse(nums, 1, len(nums)-1)

	if stolenFirst > stolenEnd {
		return stolenFirst
	}

	return stolenEnd
}

func robHouse(nums []int, startIdx, endIdx int) int {
	stolenPre, notStolenPre := 0, 0

	for idx := startIdx; idx <= endIdx; idx += 1 {
		// 偷当前: notStolenPre+nums[idx]
		// 不偷当前: max(stolenPre, notStolenPre)
		if stolenPre > notStolenPre {
			stolenPre, notStolenPre = notStolenPre+nums[idx], stolenPre
		} else {
			stolenPre, notStolenPre = notStolenPre+nums[idx], notStolenPre
		}
	}

	if stolenPre > notStolenPre {
		return stolenPre
	}

	return notStolenPre
}

// @lc code=end

