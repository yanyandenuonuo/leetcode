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
	// solution1: 动态规划，自下而上
	/**
	if amount == 0 {
		return 0
	}
	coinMap := make(map[int]int, amount+1)
	for subAmount := 1; subAmount <= amount; subAmount += 1 {
		for _, coin := range coins {
			if subAmount-coin < 0 {
				continue
			}

			subCoinCount := coinMap[subAmount-coin]
			if subAmount-coin != 0 && subCoinCount == 0 {
				continue
			}

			coinCount := coinMap[subAmount]

			if coinCount == 0 || coinCount > subCoinCount+1 {
				coinMap[subAmount] = subCoinCount + 1
			}
		}
	}

	if coinMap[amount] > 0 {
		return coinMap[amount]
	}
	return -1
	*/

	// solution1: 动态规划，自上而下
	return coinChangeWithMemory(coins, amount, make(map[int]int, amount))
}

func coinChangeWithMemory(coins []int, amount int, memory map[int]int) int {
	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	} else if memory[amount] != 0 {
		return memory[amount]
	} else {
		for _, coin := range coins {
			if amount-coin < 0 {
				continue
			}

			subAmount := amount - coin
			subAmountCount := memory[subAmount]

			if subAmountCount == -1 {
				continue
			}

			if subAmount != 0 && subAmountCount == 0 {
				subAmountCount = coinChangeWithMemory(coins, subAmount, memory)
				if subAmountCount == -1 {
					memory[subAmount] = -1
					continue
				}
			}

			if memory[amount] == 0 || memory[amount] > subAmountCount+1 {
				memory[amount] = subAmountCount + 1
			}
		}
	}
	if memory[amount] > 0 {
		return memory[amount]
	}

	return -1
}

// @lc code=end

