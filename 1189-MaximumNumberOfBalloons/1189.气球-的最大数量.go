/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	// solution: 统计balloon的数量，只有a, b, l, o, n，可以将l和o的数量除2，然后取这几个字符出现的最小值即可
	charCounter := make([]int, 26)
	for _, runeChar := range text {
		switch runeChar {
		case 'b', 'a', 'l', 'o', 'n':
			charCounter[runeChar-'a'] += 1
		}
	}
	charCounter['l'-'a'] /= 2
	charCounter['o'-'a'] /= 2
	minCounter := len(text)
	for _, runeChar := range "balon" {
		if charCounter[runeChar-'a'] < minCounter {
			minCounter = charCounter[runeChar-'a']
		}
	}
	return minCounter
}

// @lc code=end

