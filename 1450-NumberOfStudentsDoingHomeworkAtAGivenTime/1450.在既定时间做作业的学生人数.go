/*
 * @lc app=leetcode.cn id=1450 lang=golang
 *
 * [1450] 在既定时间做作业的学生人数
 */

// @lc code=start
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	counter := 0
	for _, currentTime := range startTime {
		if currentTime <= queryTime {
			counter += 1
		}
	}

	for _, currentTime := range endTime {
		if currentTime < queryTime {
			counter -= 1
		}
	}

	return counter
}

// @lc code=end

