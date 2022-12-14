/*
 * @lc app=leetcode.cn id=2295 lang=golang
 *
 * [2295] 替换数组中的元素
 */

// @lc code=start
func arrayChange(nums []int, operations [][]int) []int {
	idxMap := make(map[int]int, len(nums))
	for idx, val := range nums {
		idxMap[val] = idx
	}
	for _, operation := range operations {
		idx := idxMap[operation[0]]
		nums[idx] = operation[1]

		delete(idxMap, operation[0])
		idxMap[operation[1]] = idx
	}
	return nums
}

// @lc code=end

