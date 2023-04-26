/*
 * @lc app=leetcode.cn id=1035 lang=golang
 *
 * [1035] 不相交的线
 */

// @lc code=start
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	// solution1: 动态规划，实质为求2个数组的最长公共子序列问题
	//						num2-0	num2-1	num2-2
	//				num1-0	  0		  0		  0
	//				num1-1	  0
	//				num1-2	  0

	dp := make([][]int, len(nums1)+1)
	for idx1 := 0; idx1 < len(nums1)+1; idx1 += 1 {
		dp[idx1] = make([]int, len(nums2)+1)
	}
	for idx1 := 0; idx1 < len(nums1); idx1 += 1 {
		for idx2 := 0; idx2 < len(nums2); idx2 += 1 {
			if nums1[idx1] == nums2[idx2] {
				dp[idx1+1][idx2+1] = dp[idx1][idx2] + 1
			} else {
				dp[idx1+1][idx2+1] = dp[idx1+1][idx2]
				if dp[idx1][idx2+1] > dp[idx1+1][idx2+1] {
					dp[idx1+1][idx2+1] = dp[idx1][idx2+1]
				}
			}
		}
	}

	return dp[len(nums1)][len(nums2)]
}

// @lc code=end

