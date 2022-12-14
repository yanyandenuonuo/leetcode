/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	dp[0] = make([]bool, len(p)+1)
	dp[0][0] = true
	for pIdx := 0; pIdx < len(p); pIdx += 1 {
		if p[pIdx] == '*' && dp[0][pIdx-1] {
			dp[0][pIdx+1] = true
		}
	}
	for sIdx := 0; sIdx < len(s); sIdx += 1 {
		dp[sIdx+1] = make([]bool, len(p)+1)
		for pIdx := 0; pIdx < len(p); pIdx += 1 {
			if p[pIdx] == '.' || p[pIdx] == s[sIdx] {
				dp[sIdx+1][pIdx+1] = dp[sIdx][pIdx]
			} else if p[pIdx] == '*' {
				if p[pIdx-1] == '.' || p[pIdx-1] == s[sIdx] {
					// 不匹配 || 匹配1次 || 匹配多次
					dp[sIdx+1][pIdx+1] = dp[sIdx+1][pIdx-1] || dp[sIdx+1][pIdx] || dp[sIdx][pIdx+1]
				} else {
					// 不匹配
					dp[sIdx+1][pIdx+1] = dp[sIdx+1][pIdx-1]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

// @lc code=end

