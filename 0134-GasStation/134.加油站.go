/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	// solution: 从第i个加油站出发：
	//				到达i，返回i
	//				到达j，从j+1出发继续判断
	//				从第i个站出发，说明到达i-j其中任一站剩余gas一定大于等于0，如果大于等于0都不能到达j+1，那从这里任一站出发一定无法到达j+1
	for idx := 0; idx < len(gas); {
		gasTank := gas[idx] - cost[idx]
		nextIdx := (idx + 1) % len(gas)
		for gasTank >= 0 {
			if nextIdx == idx {
				return idx
			}
			gasTank += gas[nextIdx] - cost[nextIdx]
			nextIdx = (nextIdx + 1) % len(gas)
		}
		// 重复遍历场景
		if nextIdx <= idx {
			return -1
		}
		idx = nextIdx
	}
	return -1
}

// @lc code=end

