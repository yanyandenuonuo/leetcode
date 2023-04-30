/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]bool, len(nums1))
	for _, num := range nums1 {
		numMap[num] = true
	}
	targetMap := make(map[int]bool, len(numMap))
	for _, num := range nums2 {
		if numMap[num] {
			targetMap[num] = true
		}
	}
	res := make([]int, 0, len(targetMap))
	for num := range targetMap {
		res = append(res, num)
	}
	return res
}

// @lc code=end

