/*
 * @lc app=leetcode.cn id=914 lang=golang
 *
 * [914] 卡牌分组
 */

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
	// solution: 找cardCount最大公约数
	cardCount := make(map[int]int, len(deck)/2)

	for _, card := range deck {
		cardCount[card] += 1
	}

	gcd := 0

	for _, val := range cardCount {
		if gcd == 0 {
			gcd = val
			continue
		}

		gcd = findGCD(val, gcd)
	}

	return gcd >= 2
}

func findGCD(maxVal, minVal int) int {
	for minVal != 0 {
		minVal, maxVal = maxVal%minVal, minVal
	}

	return maxVal
}

// @lc code=end

