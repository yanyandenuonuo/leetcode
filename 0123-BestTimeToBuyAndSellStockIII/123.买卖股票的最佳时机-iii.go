/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0

	for idx := 1; idx < len(prices); idx += 1 {
		buy1 = maxValue(buy1, -prices[idx])
		sell1 = maxValue(sell1, buy1+prices[idx])

		buy2 = maxValue(buy2, sell1-prices[idx])
		sell2 = maxValue(sell2, buy2+prices[idx])
	}

	return sell2
}

func maxValue(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// @lc code=end

