/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	// 动态规划+3指针
	uglyNum := make([]int, n)
	uglyNum[0] = 1
	for uglyIdx, idx2, idx3, idx5 := 1, 0, 0, 0; uglyIdx < n; uglyIdx += 1 {
		num2 := uglyNum[idx2] * 2
		num3 := uglyNum[idx3] * 3
		num5 := uglyNum[idx5] * 5

		targetVal := 0
		if num2 <= num3 && num2 <= num5 {
			targetVal = num2
		} else if num3 <= num2 && num3 <= num5 {
			targetVal = num3
		} else if num5 <= num2 && num5 <= num3 {
			targetVal = num5
		}

		uglyNum[uglyIdx] = targetVal

		if targetVal == num2 {
			idx2 += 1
		}
		if targetVal == num3 {
			idx3 += 1
		}
		if targetVal == num5 {
			idx5 += 1
		}
	}
	return uglyNum[n-1]
}

// @lc code=end

