/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(jewels string, stones string) int {
	// solution1: hash
	// solution2: 计数
	counterMap := make(map[rune]int, len(jewels))
	for _, stone := range stones {
		counterMap[stone] += 1
	}

	counter := 0
	for _, jewel := range jewels {
		counter += counterMap[jewel]
	}

	return counter
}

// @lc code=end

