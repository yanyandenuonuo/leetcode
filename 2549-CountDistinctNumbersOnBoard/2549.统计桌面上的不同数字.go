/*
 * @lc app=leetcode.cn id=2549 lang=golang
 *
 * [2549] 统计桌面上的不同数字
 */

// @lc code=start
func distinctIntegers(n int) int {
	// solution1: 模拟
	/**
	numMap := map[int]struct{}{
		n: struct{}{},
	}

	for day := 0; day <= 1_000_000_000; day += 1 {
		dayNumMap := make(map[int]struct{}, 0)

		for idx := 1; idx <= n; idx += 1 {
			for num := range numMap {
				if num%idx == 1 {
					dayNumMap[idx] = struct{}{}
				}
			}
		}

		for num := range dayNumMap {
			numMap[num] = struct{}{}
		}
	}

	return len(numMap)
	*/

	// solution2: x-1必出现在桌子上，所以时间足够时2-n均会出现在桌子上（1%1为0）
	if n == 1 {
		return 1
	}

	return n - 1
}

// @lc code=end

