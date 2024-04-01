/*
 * @lc app=leetcode.cn id=1431 lang=golang
 *
 * [1431] 拥有最多糖果的孩子
 */

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	// solution: 找最大值
	maxVal := 0
	for _, candie := range candies {
		if candie > maxVal {
			maxVal = candie
		}
	}

	resList := make([]bool, len(candies))

	for idx, candie := range candies {
		resList[idx] = candie+extraCandies >= maxVal
	}

	return resList
}

// @lc code=end

