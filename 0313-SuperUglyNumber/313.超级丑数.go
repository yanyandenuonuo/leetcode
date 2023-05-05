/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	// solution: 动态规划
	pointers := make([]int, len(primes))
	nums := make([]int, len(primes))
	for idx := 0; idx < len(primes); idx += 1 {
		nums[idx] = 1
	}

	dp := make([]int, n+1)

	for idx := 1; idx <= n; idx += 1 {
		minNum, minIdxList := getMinNumAndIdx(nums)
		dp[idx] = minNum

		for _, minIdx := range minIdxList {
			pointers[minIdx] += 1
			nums[minIdx] = dp[pointers[minIdx]] * primes[minIdx]
		}
	}

	return dp[n]
}

func getMinNumAndIdx(nums []int) (int, []int) {
	minIdxList := []int{0}
	for idx := 1; idx < len(nums); idx += 1 {
		if nums[idx] < nums[minIdxList[0]] {
			minIdxList = []int{idx}
		} else if nums[idx] == nums[minIdxList[0]] {
			minIdxList = append(minIdxList, idx)
		}
	}
	return nums[minIdxList[0]], minIdxList
}

// @lc code=end

