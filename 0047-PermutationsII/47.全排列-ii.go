/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	// solution: 搜索回溯
	if len(nums) == 0 {
		return [][]int{}
	} else if len(nums) == 1 {
		return [][]int{nums}
	}

	resList := make([][]int, 0)

	// 重复的数字只需要遍历一次
	memory := make(map[int]bool, len(nums))

	for idx := 0; idx < len(nums); idx += 1 {
		if memory[nums[idx]] {
			continue
		}
		memory[nums[idx]] = true

		subNums := make([]int, len(nums)-1)
		copy(subNums, nums[:idx])
		copy(subNums[idx:], nums[idx+1:])

		subResList := permuteUnique(subNums)
		for _, subRes := range subResList {
			resList = append(resList, append([]int{nums[idx]}, subRes...))
		}
	}

	return resList
}

// @lc code=end

