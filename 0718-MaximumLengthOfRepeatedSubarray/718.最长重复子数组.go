/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	// 转化为二维数组进行比较, 如下图
	// \	-	1	2	3	4	5
	// -	0	0	0	0	0	0
	// 0	0	0	0	0	0	0
	// 2	0	0	1	0	0	0
	// 3	0	0	0	2	0	0
	// 4	0	0	0	0	3	0
	// 6	0	0	0	0	0	0
	dp := make([][]int, len(nums1)+1)
	dp[0] = make([]int, len(nums2)+1)

	maxLength := 0
	for idx1 := 1; idx1 <= len(nums1); idx1 += 1 {
		dp[idx1] = make([]int, len(nums2)+1)
		for idx2 := 1; idx2 <= len(nums2); idx2 += 1 {
			if nums1[idx1-1] == nums2[idx2-1] {
				dp[idx1][idx2] = dp[idx1-1][idx2-1] + 1
			}

			if dp[idx1][idx2] > maxLength {
				maxLength = dp[idx1][idx2]
			}
		}
	}
	return maxLength
}

// @lc code=end

