/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	// solution1: 使用hash检测是不是存在环
	//			  无论是多大的数，最终都会降维到3位数内，所以可以直接推理
	//			  9 	=> 9*9 		=> 81
	//			  99 	=> 9*9 * 2	=> 162
	//			  999 	=> 9*9 * 3	=> 243
	//			  9999 	=> 9*9 * 4	=> 324
	//			  99999	=> 9*9 * 5	=> 405
	/**
	numMap := make(map[int]bool)
	for n = getNumSum(n); n != 1 && !numMap[n]; n = getNumSum(n) {
		numMap[n] = true
	}
	return n == 1
	*/

	// solution2: 使用快慢指针，若存在循环，则快慢指针会相等，否则快指针会等于1
	slowN := getNumSum(n)
	fastN := getNumSum(getNumSum(n))
	for slowN != fastN && fastN != 1 {
		slowN = getNumSum(slowN)
		fastN = getNumSum(getNumSum(fastN))
	}
	return fastN == 1
}

func getNumSum(num int) int {
	sum := 0
	for ; num > 0; num /= 10 {
		sum += (num % 10) * (num % 10)
	}
	return sum
}

// @lc code=end

