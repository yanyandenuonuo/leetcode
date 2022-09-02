/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (36.16%)
 * Likes:    5219
 * Dislikes: 158
 * Total Accepted:    492.3K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given coins of different denominations and a total amount of money
 * amount. Write a function to compute the fewest number of coins that you need
 * to make up that amount. If that amount of money cannot be made up by any
 * combination of the coins, return -1.
 * 
 * You may assume that you have an infinite number of each kind of coin.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: coins = [1,2,5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: coins = [2], amount = 3
 * Output: -1
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: coins = [1], amount = 0
 * Output: 0
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: coins = [1], amount = 1
 * Output: 1
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: coins = [1], amount = 2
 * Output: 2
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 * 
 * 
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	coinMap := make(map[int]int, amount)
	coinMap[0] = 0
	for idx := 1; idx <= amount; idx += 1 {
		for _, coin := range coins {
			if idx-coin < 0 {
				continue
			}
			if _, isExist := coinMap[idx]; !isExist {
				coinMap[idx] = -1
			}

			coinCount := coinMap[idx]

			lastCoinCount, isExist := coinMap[idx-coin]
			if !isExist || lastCoinCount == -1 {
				continue
			}

			if coinCount == -1 || coinCount > lastCoinCount+1 {
				coinMap[idx] = lastCoinCount + 1
			}
		}
	}
	if coinCount, isExist := coinMap[amount]; isExist {
		return coinCount
	}
	return -1
}
// @lc code=end

