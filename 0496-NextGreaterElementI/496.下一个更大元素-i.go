/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// solution1: 遍历
	// solution2: 倒序遍历，nums2，利用栈获取每个元素下一个更大元素，使用map记录，然后遍历nums1即可
	nextMap := make(map[int]int, len(nums2))
	numStack := make([]int, 0, len(nums2))
	for idx := len(nums2) - 1; idx >= 0; idx -= 1 {
		for len(numStack) > 0 && numStack[len(numStack)-1] <= nums2[idx] {
			numStack = numStack[:len(numStack)-1]
		}
		if len(numStack) > 0 {
			nextMap[nums2[idx]] = numStack[len(numStack)-1]
		} else {
			nextMap[nums2[idx]] = -1
		}
		numStack = append(numStack, nums2[idx])
	}

	res := make([]int, len(nums1))
	for idx, num := range nums1 {
		res[idx] = nextMap[num]
	}

	return res
}

// @lc code=end

