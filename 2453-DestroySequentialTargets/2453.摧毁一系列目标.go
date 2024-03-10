/*
 * @lc app=leetcode.cn id=2453 lang=golang
 *
 * [2453] 摧毁一系列目标
 */

// @lc code=start
func destroyTargets(nums []int, space int) int {
	// solution1: 暴力
	// solution2: 按nums[i]%space分组计数，计数最大值mod的最小值即为对应数值
	modCounterMap, modMinValMap, maxCounter := make(map[int]int, 0), make(map[int]int, 0), 0
	for _, val := range nums {
		mod := val % space

		modCounterMap[mod] += 1

		if modCounterMap[mod] > maxCounter {
			maxCounter = modCounterMap[mod]
		}

		if _, isExist := modMinValMap[mod]; !isExist || val < modMinValMap[mod] {
			modMinValMap[mod] = val
		}
	}

	minVal := 1<<31 - 1
	for mod := range modCounterMap {
		if modCounterMap[mod] < maxCounter {
			continue
		}

		if modMinValMap[mod] < minVal {
			minVal = modMinValMap[mod]
		}
	}

	return minVal
}

// @lc code=end

