/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	for leftIdx, rightIdx := 0, (num+1)/2; leftIdx <= rightIdx; {
		midNum := leftIdx + (rightIdx-leftIdx)/2
		productNum := midNum * midNum
		if productNum > num {
			rightIdx = midNum - 1
		} else if productNum < num {
			leftIdx = midNum + 1
		} else {
			return true
		}
	}

	return false
}

// @lc code=end

