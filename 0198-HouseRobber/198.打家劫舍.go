/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	robMemory := make(map[int]int, len(nums))
	return whoIsLucky(nums, 0, robMemory)
}

func whoIsLucky(nums []int, idx int, robMemory map[int]int) int {
	if robMoney, isSet := robMemory[idx]; isSet {
		return robMoney
	}

	choiceNums := nums[idx:]
	isLucky := 0
	notLucky := 0
	if len(choiceNums) == 0 {
		robMemory[idx] = 0
		return 0
	} else if len(choiceNums) == 1 {
		robMemory[idx] = choiceNums[0]
		return choiceNums[0]
	} else if len(choiceNums) == 2 {
		isLucky = choiceNums[1]
		notLucky = choiceNums[0]
	} else {
		isLucky += whoIsLucky(nums, idx+1, robMemory)
		notLucky = choiceNums[0] + whoIsLucky(nums, idx+2, robMemory)
	}

	if isLucky > notLucky {
		robMemory[idx] = isLucky

		return isLucky
	}

	robMemory[idx] = notLucky

	return notLucky
}

// @lc code=end

