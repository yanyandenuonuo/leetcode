/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 */

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	// solution: 逐个校验
	res := make([]int, 0)
	for num := left; num <= right; num += 1 {
		if isSelfDividingNumber(num) {
			res = append(res, num)
		}
	}

	return res
}

func isSelfDividingNumber(num int) bool {
	for nextNum := num; nextNum > 0; nextNum /= 10 {
		if nextNum%10 == 0 || num%(nextNum%10) > 0 {
			return false
		}
	}

	return true
}

// @lc code=end

