/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	return backTracker(1, n, k)
}

func backTracker(start, end int, k int) [][]int {
	resultList := make([][]int, 0)
	for idx := start; idx <= end; idx += 1 {
		if k == 1 {
			resultList = append(resultList, []int{idx})
		} else {
			subResultList := backTracker(idx+1, end, k-1)
			for resultIdx := range subResultList {
				resultList = append(resultList, append([]int{idx}, subResultList[resultIdx]...))
			}
		}
	}

	return resultList
}

// @lc code=end

