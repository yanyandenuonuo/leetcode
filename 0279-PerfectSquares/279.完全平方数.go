/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	// solution1: 枚举，超时
	/**
	numList := make([]int, 0)
	for idx := 1; true; idx += 1 {
		num := idx * idx

		if num > n {
			break
		}

		numList = append(numList, num)
	}

	minCounter := 1<<31 - 1

	var findSum func(idx, sum, counter int)
	findSum = func(idx, sum, counter int) {
		if sum < 0 {
			return
		}

		if sum == 0 && counter < minCounter {
			minCounter = counter

			return
		}

		for idx >= 0 {
			findSum(idx, sum-numList[idx], counter+1)

			idx -= 1
		}

		return
	}

	findSum(len(numList)-1, n, 0)

	return minCounter
	*/

	// solution2: 动态规划，dp(i) = 1+ min(dp(i-x*x))，枚举i和x
	dp := make([]int, n+1)
	for idx := 1; idx <= n; idx += 1 {
		dp[idx] = 1<<31 - 1

		for x := 1; x*x <= idx; x += 1 {
			if dp[idx] > 1+dp[idx-x*x] {
				dp[idx] = 1 + dp[idx-x*x]
			}
		}
	}

	return dp[n]

	// solution3: 四平方和定理
	//			  任意一个正整数都可以被表示为至多四个正整数的平方和
	//			  当且仅当n!=4^k*(8m+7)时，n可以被表示为至多三个正整数的平方和
	//			  1 => n为完全平方数
	//			  2 => n = x^2  + y^2
	//			  4 => n = 4^k * (8m+7)
	//			  3 => 其余
}

// @lc code=end

