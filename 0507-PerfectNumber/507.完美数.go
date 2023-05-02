/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 */

// @lc code=start
func checkPerfectNumber(num int) bool {
	// solution1: 枚举，可以枚举到num-1，也可以剪枝只枚举到sqrt(num)，然后将num/idx的值同步加到sum
	sum := 0
	for idx := 1; idx*idx <= num; idx += 1 {
		if num%idx != 0 || idx == num {
			continue
		}
		sum += idx
		if num/idx != idx && num/idx != num {
			sum += num / idx
		}
	}
	return sum == num

	// solution2: 数学，欧几里得-欧拉定理，10^8以内的完美数一共有6, 28, 496, 8128, 33550336
	/**
	return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
	*/
}

// @lc code=end

