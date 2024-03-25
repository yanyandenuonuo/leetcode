/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 */

// @lc code=start
func bulbSwitch(n int) int {
	// solution1: 模拟
	// solution2: 找约数，初始关灯，则n轮后有奇数约数的灯是开着的
	//			  假设x存在约数k，则x必存在约数x/k，即当k!=x/k时x必有偶数个约数
	//			  k==x/k= > x==k^2时x称为完全平方数
	//			  1-n之间的完全平方数为开根号n向下取整
	//			  1^2 < n, 2^2 < n 3^2 < n, 4^2 < n, ..., i^2 <= n, (i+1)^2 > n
	counter := 0
	for idx := 1; idx*idx <= n; idx += 1 {
		counter += 1
	}
	// 上述等价与math.Sqrt(float64(n))，考虑浮点数精度影响也等价于math.Sqrt(float64(n)+0.5)

	return counter
}

// @lc code=end

