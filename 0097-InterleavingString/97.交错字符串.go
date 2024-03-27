/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	// warning: 无法使用双指针，因为存在需要交换顺序的case，如s1="aabcc", s2="dbbca", s3="aadbbcbcac"
	// solution: 动态规划，dp[i][j]表示s3[0:i+j+1]能由s1[:i+1]和s2[:j+1]组成
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)

	for i := 0; i <= len(s1); i += 1 {
		dp[i] = make([]bool, len(s2)+1)

		for j := 0; j <= len(s2); j += 1 {
			if i == 0 && j == 0 {
				dp[0][0] = true
				continue
			}

			if (i > 0 && dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (j > 0 && dp[i][j-1] && s2[j-1] == s3[i+j-1]) {
				dp[i][j] = true
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

// @lc code=end

