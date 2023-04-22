/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return [][]int{}
	}

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	memory := make([][]bool, len(heights))
	for idx := range memory {
		memory[idx] = make([]bool, len(heights[idx]))
	}

	var dfs func(rowIdx, columnIdx int, toPacific, toAtlantic bool) (bool, bool)
	dfs = func(rowIdx, columnIdx int, toPacific, toAtlantic bool) (bool, bool) {
		if rowIdx == 0 || columnIdx == 0 {
			toPacific = true
		}
		if rowIdx == len(heights)-1 || columnIdx == len(heights[0])-1 {
			toAtlantic = true
		}

		if toPacific && toAtlantic {
			return toPacific, toAtlantic
		}

		memory[rowIdx][columnIdx] = true
		defer func() {
			memory[rowIdx][columnIdx] = false
		}()
		for _, direction := range directions {
			nextRowIdx := rowIdx + direction[0]
			nextColumnIdx := columnIdx + direction[1]
			if nextRowIdx < 0 || nextRowIdx >= len(heights) || nextColumnIdx < 0 || nextColumnIdx >= len(heights[0]) ||
				memory[nextRowIdx][nextColumnIdx] || heights[rowIdx][columnIdx] < heights[nextRowIdx][nextColumnIdx] {
				continue
			}

			toPacific, toAtlantic = dfs(nextRowIdx, nextColumnIdx, toPacific, toAtlantic)
			if toPacific && toAtlantic {
				break
			}
		}
		return toPacific, toAtlantic
	}

	res := make([][]int, 0)
	for rowIdx := 0; rowIdx < len(heights); rowIdx += 1 {
		for columnIdx := 0; columnIdx < len(heights[rowIdx]); columnIdx += 1 {
			toPacific, toAtlantic := dfs(rowIdx, columnIdx, false, false)
			if toPacific && toAtlantic {
				res = append(res, []int{rowIdx, columnIdx})
			}
		}
	}

	return res
}

// @lc code=end

