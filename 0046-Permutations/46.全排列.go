/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	return backTracker(nums, len(nums))
}

func backTracker(nums []int, k int) [][]int {
	resultList := make([][]int, 0)
	for idx := range nums {
		if k == 1 {
			resultList = append(resultList, []int{nums[idx]})
		} else {
			subResultList := backTracker(append(nums[:idx], nums[idx+1:]...), k-1)
			for resultIdx := range subResultList {
				resultList = append(resultList, append([]int{nums[idx]}, subResultList[resultIdx]...))
			}
		}
	}

	return resultList
}

// @lc code=end

