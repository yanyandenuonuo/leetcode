/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	m -= 1
	n -= 1
	for idx := m + n + 1; m >= 0 && n >= 0; idx -= 1 {
		if nums1[m] > nums2[n] {
			nums1[idx] = nums1[m]
			m -= 1
		} else {
			nums1[idx] = nums2[n]
			n -= 1
		}
	}
	for n >= 0 {
		nums1[n] = nums2[n]
		n -= 1
	}
}

// @lc code=end

