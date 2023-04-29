/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	// solution1: hash表
	/**
	numMap := make(map[int]int, len(nums))
	for idx, val := range nums {
		if preIdx, isExist := numMap[val]; isExist {
			if idx-preIdx <= k {
				return true
			}
		}
		numMap[val] = idx
	}
	return false
	*/
	// solution2: 滑动窗口+hash表
	numMap := make(map[int]bool, k+1)
	for idx, val := range nums {
		if idx > k {
			delete(numMap, nums[idx-k-1])
		}
		if numMap[val] {
			return true
		}
		numMap[val] = true
	}
	return false
}

// @lc code=end

