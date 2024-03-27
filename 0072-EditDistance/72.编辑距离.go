/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	// solution1: 递归模拟
	// solution2: 动态规划 dp[i][j]表示word1[:i]与word2[:j]之间的编辑距离
	//			  对其中一个单词进行插入等于对另一个单词进行删除
	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	dp := make([][]int, len(word1)+1)

	for i := 0; i <= len(word1); i += 1 {
		dp[i] = make([]int, len(word2)+1)

		for j := 0; j <= len(word2); j += 1 {
			if i == 0 || j == 0 {
				dp[i][j] = i + j
				continue
			}

			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			// 插入
			dp[i][j] = 1 + dp[i-1][j]

			// 删除
			if 1+dp[i][j-1] < dp[i][j] {
				dp[i][j] = 1 + dp[i][j-1]
			}

			// 替换
			if 1+dp[i-1][j-1] < dp[i][j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

// @lc code=end

