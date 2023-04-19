/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// solution1: 动态规划
	wordMap := make(map[string]bool, len(wordDict))

	for _, word := range wordDict {
		wordMap[word] = true
	}

	dp := make(map[int]bool, len(s)+1)
	dp[0] = true
	for leftIdx := 0; leftIdx < len(s); leftIdx += 1 {
		if !dp[leftIdx] {
			continue
		}
		for rightIdx := leftIdx; rightIdx < len(s); rightIdx += 1 {
			if wordMap[s[leftIdx:rightIdx+1]] && dp[leftIdx] {
				dp[rightIdx+1] = true
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end

