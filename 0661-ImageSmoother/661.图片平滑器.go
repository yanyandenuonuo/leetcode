/*
 * @lc app=leetcode.cn id=661 lang=golang
 *
 * [661] 图片平滑器
 */

// @lc code=start
func imageSmoother(img [][]int) [][]int {
	// solution: 模拟
	res := make([][]int, len(img))
	for rowIdx := 0; rowIdx < len(img); rowIdx += 1 {
		res[rowIdx] = make([]int, len(img[rowIdx]))

		for columnIdx := 0; columnIdx < len(img[rowIdx]); columnIdx += 1 {
			counter, sum := 0, 0
			for filterRowIdx := rowIdx - 1; filterRowIdx <= rowIdx+1; filterRowIdx += 1 {
				if filterRowIdx < 0 || filterRowIdx >= len(img) {
					continue
				}

				for filterColumnIdx := columnIdx - 1; filterColumnIdx <= columnIdx+1; filterColumnIdx += 1 {
					if filterColumnIdx < 0 || filterColumnIdx >= len(img[rowIdx]) {
						continue
					}

					counter += 1
					sum += img[filterRowIdx][filterColumnIdx]
				}
			}

			res[rowIdx][columnIdx] = sum / counter
		}
	}

	return res
}

// @lc code=end

