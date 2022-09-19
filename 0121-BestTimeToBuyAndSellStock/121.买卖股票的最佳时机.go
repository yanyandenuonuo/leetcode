/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// solution1: 暴力求解，2层循环
	// solution2: 求极值，最小值及最大差值
	minValue, maxDiff := prices[0], 0

	for idx := 1; idx < len(prices); idx += 1 {
		if prices[idx] < minValue {
			minValue = prices[idx]
			continue
		}
		diff := prices[idx] - minValue
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	if maxDiff < 0 {
		maxDiff = 0
	}
	return maxDiff

	// solution3: 变换数组，存储相邻两个值的差值，然后求最大子序列和
	/*
		subSum, maxSubSum := 0, 0
		for idx := 1; idx < len(prices); idx += 1 {
			subSum += (prices[idx] - prices[idx-1])
			if subSum < 0 {
				subSum = 0
			}
			if subSum > maxSubSum {
				maxSubSum = subSum
			}
		}
		return maxSubSum
	*/
}

// @lc code=end

