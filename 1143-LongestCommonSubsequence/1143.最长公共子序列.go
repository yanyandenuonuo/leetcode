/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	// 转化为二维数组进行比较, 如下图
	// \	''	a	b	c	d	e
	// ''	0	0	0	0	0	0
	// a	0	1	1	1	1	1
	// c	0	1	1	2	2	2
	// f	0	1	1	2	2	2
	// d	0	1	1	2	3	3

	dp := make([][]int, len(text1)+1)
	dp[0] = make([]int, len(text2)+1)

	for idx1 := 1; idx1 <= len(text1); idx1 += 1 {
		dp[idx1] = make([]int, len(text2)+1)
		for idx2 := 1; idx2 <= len(text2); idx2 += 1 {
			if text1[idx1-1] == text2[idx2-1] {
				dp[idx1][idx2] = dp[idx1-1][idx2-1] + 1
			} else {
				if dp[idx1-1][idx2] > dp[idx1][idx2-1] {
					dp[idx1][idx2] = dp[idx1-1][idx2]
				} else {
					dp[idx1][idx2] = dp[idx1][idx2-1]
				}
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// @lc code=end

