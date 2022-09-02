/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lengthSum := len(nums1) + len(nums2)
	midNum1 := 0
	midNum2 := 0

	num1Idx := 0
	num2Idx := 0

	for idx := 0; idx <= lengthSum/2; idx += 1 {
		midNum1 = midNum2

		if num1Idx >= len(nums1) {
			midNum2 = nums2[num2Idx]
			num2Idx += 1
			continue
		}

		if num2Idx >= len(nums2) {
			midNum2 = nums1[num1Idx]
			num1Idx += 1
			continue
		}

		if nums1[num1Idx] < nums2[num2Idx] {
			midNum2 = nums1[num1Idx]
			num1Idx += 1
		} else {
			midNum2 = nums2[num2Idx]
			num2Idx += 1
		}
	}

	if lengthSum%2 == 0 {
		return (float64(midNum1) + float64(midNum2)) / 2
	}

	return float64(midNum2)
}

// @lc code=end

