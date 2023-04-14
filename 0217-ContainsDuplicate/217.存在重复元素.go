/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	valMap := make(map[int]bool, len(nums))
	for _, val := range nums {
		if valMap[val] {
			return true
		}
		valMap[val] = true
	}
	return false
}

// @lc code=end

