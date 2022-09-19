/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	minCostSum := make([]int, len(cost))
	minCostSum[0] = cost[0]
	minCostSum[1] = cost[1]

	for idx := 2; idx < len(cost); idx += 1 {
		minCostSum[idx] = minValue(minCostSum[idx-2], minCostSum[idx-1]) + cost[idx]
	}
	return minValue(minCostSum[len(cost)-2], minCostSum[len(cost)-1])
}

func minValue(preTwo, preOne int) int {
	if preTwo < preOne {
		return preTwo
	}
	return preOne
}

// @lc code=end

