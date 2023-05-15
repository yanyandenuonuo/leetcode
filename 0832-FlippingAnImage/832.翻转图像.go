/*
 * @lc app=leetcode.cn id=832 lang=golang
 *
 * [832] 翻转图像
 */

// @lc code=start
func flipAndInvertImage(image [][]int) [][]int {
	// solution: 模拟
	for rowIdx := 0; rowIdx < len(image); rowIdx += 1 {
		for leftIdx, rightIdx := 0, len(image[rowIdx])-1; leftIdx <= rightIdx; {
			image[rowIdx][leftIdx], image[rowIdx][rightIdx] = image[rowIdx][rightIdx], image[rowIdx][leftIdx]
			if image[rowIdx][leftIdx] == 0 {
				image[rowIdx][leftIdx] = 1
			} else {
				image[rowIdx][leftIdx] = 0
			}

			if leftIdx == rightIdx {
				break
			}

			if image[rowIdx][rightIdx] == 0 {
				image[rowIdx][rightIdx] = 1
			} else {
				image[rowIdx][rightIdx] = 0
			}

			leftIdx += 1
			rightIdx -= 1
		}
	}

	return image
}

// @lc code=end

