/*
 * @lc app=leetcode.cn id=1668 lang=golang
 *
 * [1668] 最大重复子字符串
 */

// @lc code=start
func maxRepeating(sequence string, word string) int {
	// solution1: 枚举+动态规划
	repeatCount, maxRepeat := make([]int, len(sequence)+1), 0

	for idx := len(word); idx <= len(sequence); idx += 1 {
		if sequence[idx-len(word):idx] != word {
			continue
		}

		repeatCount[idx] = repeatCount[idx-len(word)] + 1

		if repeatCount[idx] > maxRepeat {
			maxRepeat = repeatCount[idx]
		}
	}

	return maxRepeat

	// solution2: kmp+动态规划
}

// @lc code=end

