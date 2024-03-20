/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
func prefixesDivBy5(nums []int) []bool {
	// solution1: 直接计算即可，数组太长会溢出，每次保留除5的余数即可
	//			  nums[idx]%5 = (nums[idx-1] << 1 + nums[idx]) % 5
	//			  nums[idx]%5 = (nums[idx-1] << 1) % 5 + nums[idx] % 5

	resList := make([]bool, len(nums))
	remainderNum := 0
	for idx, val := range nums {
		remainderNum = (remainderNum<<1 | val) % 5

		resList[idx] = remainderNum == 0
	}

	return resList
}

// @lc code=end

