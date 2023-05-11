/*
 * @lc app=leetcode.cn id=575 lang=golang
 *
 * [575] 分糖果
 */

// @lc code=start
func distributeCandies(candyType []int) int {
	// solution: hash
	typeMap := make(map[int]bool, 0)
	for _, candy := range candyType {
		typeMap[candy] = true
	}

	if len(typeMap) < len(candyType)/2 {
		return len(typeMap)
	}

	return len(candyType) / 2
}

// @lc code=end

