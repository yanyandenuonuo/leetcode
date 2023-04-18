/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	return backTracker(nums, len(nums))
}

func backTracker(nums []int, length int) [][]int {
	if length == 0 {
		return [][]int{}
	} else if length == 1 {
		return [][]int{
			{nums[0]},
		}
	}

	resultList := make([][]int, 0)
	for idx := range nums {
		subNums := make([]int, length-1)
		copy(subNums, nums[:idx])
		copy(subNums[idx:], nums[idx+1:])
		subResultList := backTracker(subNums, length-1)
		for _, subResult := range subResultList {
			resultList = append(resultList, append([]int{nums[idx]}, subResult...))
		}
	}

	return resultList
}

// @lc code=end

