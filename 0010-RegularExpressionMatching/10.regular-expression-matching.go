/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (27.01%)
 * Likes:    4826
 * Dislikes: 772
 * Total Accepted:    471.3K
 * Total Submissions: 1.7M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*' where:
 *
 *
 * '.' Matches any single character.​​​​
 * '*' Matches zero or more of the preceding element.
 *
 *
 * The matching should cover the entire input string (not partial).
 *
 *
 * Example 1:
 *
 *
 * Input: s = "aa", p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "aa", p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the preceding element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 *
 *
 * Example 3:
 *
 *
 * Input: s = "ab", p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 *
 *
 * Example 4:
 *
 *
 * Input: s = "aab", p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore,
 * it matches "aab".
 *
 *
 * Example 5:
 *
 *
 * Input: s = "mississippi", p = "mis*is*p*."
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 20
 * 0 <= p.length <= 30
 * s contains only lowercase English letters.
 * p contains only lowercase English letters, '.', and '*'.
 *
 *
 */

// @lc code=start
func isMatch(s string, p string) bool {
	// solution: 动态规划
	// dp[i+1][j+1]表示s的前i个字符与p中的前j个字符是否能够匹配
	dp := make([][]bool, len(s)+1)

	dp[0] = make([]bool, len(p)+1)
	dp[0][0] = true

	// dp[0][x]
	for pIdx := 0; pIdx < len(p); pIdx += 1 {
		// 只有x*x*x*才为true，所以p[pIdx] == * 并且 dp[0][pIdx-1]为true时dp[0][pIdx+1]才为true
		if p[pIdx] == '*' {
			dp[0][pIdx+1] = dp[0][pIdx-1]
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

