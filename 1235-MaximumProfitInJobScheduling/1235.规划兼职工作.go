/*
 * @lc app=leetcode.cn id=1235 lang=golang
 *
 * [1235] 规划兼职工作
 */

// @lc code=start
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	// solution1: 动态规划，dp[x]表示x时间节点作为结束时间的最大盈利，时间太大会超时
	/**
	if len(endTime) == 0 {
		return 0
	}

	maxEndTime := endTime[0]

	jobMap := make(map[int][][2]int, len(endTime))

	for idx := range endTime {
		if endTime[idx] > maxEndTime {
			maxEndTime = endTime[idx]
		}
		jobMap[endTime[idx]] = append(jobMap[endTime[idx]], [2]int{startTime[idx], profit[idx]})
	}

	dp := make([]int, maxEndTime+1)

	for idx := 1; idx <= maxEndTime; idx += 1 {
		dp[idx] = dp[idx-1]
		for _, job := range jobMap[idx] {
			if dp[job[0]]+job[1] > dp[idx] {
				dp[idx] = dp[job[0]] + job[1]
			}
		}
	}

	return dp[maxEndTime]
	*/

	// solution2: 动态规划，dp[x]表示job[x]结束时间时的最大盈利
	if len(endTime) == 0 {
		return 0
	}

	jobList := make([][3]int, len(endTime))

	for idx := range endTime {
		jobList[idx] = [3]int{startTime[idx], endTime[idx], profit[idx]}
	}

	// 按结束时间排序
	sort.Slice(jobList, func(i, j int) bool {
		return jobList[i][1] < jobList[j][1]
	})

	dp := make([]int, len(jobList)+1)
	for idx := 0; idx < len(jobList); idx += 1 {
		dp[idx+1] = dp[idx]

		preEndTime := sort.Search(idx, func(j int) bool {
			// preEndTime-1数组的jobList[j][1]一定小于等于jobList[idx][0]，preEndTime-1数组对应的dp为dp[preEndTime]
			return jobList[j][1] > jobList[idx][0]
		})

		if dp[preEndTime]+jobList[idx][2] > dp[idx+1] {
			dp[idx+1] = dp[preEndTime] + jobList[idx][2]
		}
	}

	return dp[len(jobList)]
}

// @lc code=end

