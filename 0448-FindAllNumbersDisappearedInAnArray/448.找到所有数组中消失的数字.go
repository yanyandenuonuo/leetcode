/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, val := range nums {
		idx := (val - 1) % n
		nums[idx] += n
	}
	res := make([]int, 0, n/2)
	for idx, val := range nums {
		if val <= n {
			res = append(res, idx+1)
		}
	}
	return res
}

// @lc code=end

