/*
 * @lc app=leetcode.cn id=495 lang=golang
 *
 * [495] 提莫攻击
 */

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) int {
	counter := 0
	for idx, freeTime := 0, 0; idx < len(timeSeries); idx += 1 {
		if freeTime <= timeSeries[idx] {
			counter += duration
		} else {
			counter += duration - (freeTime - timeSeries[idx])

		}
		freeTime = timeSeries[idx] + duration
	}
	return counter
}

// @lc code=end

