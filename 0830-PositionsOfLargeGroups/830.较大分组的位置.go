/*
 * @lc app=leetcode.cn id=830 lang=golang
 *
 * [830] 较大分组的位置
 */

// @lc code=start
func largeGroupPositions(s string) [][]int {
	resList := make([][]int, 0)

	for leftIdx := 0; leftIdx < len(s); {
		rightIdx := leftIdx

		for endIdx := leftIdx + 1; endIdx < len(s) && s[leftIdx] == s[endIdx]; endIdx += 1 {
			rightIdx = endIdx
		}

		if rightIdx-leftIdx >= 2 {
			resList = append(resList, []int{leftIdx, rightIdx})
		}

		leftIdx = rightIdx + 1
	}

	return resList
}

// @lc code=end

