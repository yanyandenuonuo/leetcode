/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		numMap[num] += 1
	}
	targetMap := make(map[int]int, len(numMap))
	for _, num := range nums2 {
		if numMap[num] > 0 {
			targetMap[num] += 1
			numMap[num] -= 1
		}
	}
	res := make([]int, 0, len(targetMap))
	for num, counter := range targetMap {
		for ; counter > 0; counter -= 1 {
			res = append(res, num)
		}
	}
	return res
}

// @lc code=end

