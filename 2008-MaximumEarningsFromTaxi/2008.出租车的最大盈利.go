/*
 * @lc app=leetcode.cn id=2008 lang=golang
 *
 * [2008] 出租车的最大盈利
 */

// @lc code=start
func maxTaxiEarnings(n int, rides [][]int) int64 {
	// solution1: 动态规划，dp[x]表示在x站点作为结束站点的最大盈利
	/**
	if n == 0 {
		return 0
	}

	stopMap := make(map[int][][2]int, len(rides))

	for _, ride := range rides {
		stopMap[ride[1]] = append(stopMap[ride[1]], [2]int{ride[0], ride[2]})
	}

	dp := make([]int64, n+1)

	for idx := 1; idx <= n; idx += 1 {
		dp[idx] = dp[idx-1]
		for _, ride := range stopMap[idx] {
			if dp[ride[0]]+int64(idx-ride[0]+ride[1]) > dp[idx] {
				dp[idx] = dp[ride[0]] + int64(idx-ride[0]+ride[1])
			}
		}
	}

	return dp[n]
	*/

	// solution2: 动态规划，dp[x]表示在rides[x]结束站点的最大盈利
	if n == 0 {
		return 0
	}

	// 按结束站点排序
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})

	dp := make([]int64, len(rides)+1)
	for idx := 0; idx < len(rides); idx += 1 {
		dp[idx+1] = dp[idx]

		preStop := sort.Search(idx, func(j int) bool {
			// preStop-1数组的rides[j][1]一定小于等于rides[idx][0]，preStop-1数组对应的dp为dp[preStop]
			return rides[j][1] > rides[idx][0]
		})

		if dp[preStop]+int64(rides[idx][1]-rides[idx][0]+rides[idx][2]) > dp[idx+1] {
			dp[idx+1] = dp[preStop] + int64(rides[idx][1]-rides[idx][0]+rides[idx][2])
		}
	}

	return dp[len(rides)]
}

// @lc code=end

