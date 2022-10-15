/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for idx := 1; idx <= n; idx += 1 {
		if idx%3 == 0 && idx%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if idx%3 == 0 {
			res = append(res, "Fizz")
		} else if idx%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(idx))
		}
	}
	return res
}

// @lc code=end

