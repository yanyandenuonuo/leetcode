/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	// solution: 模拟
	counter := 0
	for idx := 0; idx < len(flowerbed); idx += 1 {
		if flowerbed[idx] == 1 {
			continue
		}

		if (idx-1 < 0 || flowerbed[idx-1] == 0) && (idx+1 >= len(flowerbed) || flowerbed[idx+1] == 0) {
			counter += 1
			flowerbed[idx] = 1
		}
	}

	return counter >= n
}

// @lc code=end

