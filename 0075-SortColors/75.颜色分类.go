/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	for idx, zeroIdx, twoIdx := 0, 0, len(nums)-1; zeroIdx < twoIdx && idx <= twoIdx; {
		switch nums[idx] {
		case 0:
			nums[zeroIdx], nums[idx] = nums[idx], nums[zeroIdx]
			zeroIdx += 1
			idx += 1
		case 1:
			idx += 1
		case 2:
			nums[twoIdx], nums[idx] = nums[idx], nums[twoIdx]
			twoIdx -= 1
		}
	}
}

// @lc code=end

