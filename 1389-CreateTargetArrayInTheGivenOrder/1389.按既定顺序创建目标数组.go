/*
 * @lc app=leetcode.cn id=1389 lang=golang
 *
 * [1389] 按既定顺序创建目标数组
 */

// @lc code=start
func createTargetArray(nums []int, index []int) []int {
	// solution: 模拟
	target := make([]int, 0)
	for idx := 0; idx < len(nums); idx += 1 {
		if index[idx] < len(target) {
			target = append(target[:index[idx]], append([]int{nums[idx]}, target[index[idx]:]...)...)
		} else {
			target = append(target, nums[idx])
		}
	}

	return target
}

// @lc code=end

