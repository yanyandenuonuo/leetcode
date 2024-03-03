/*
 * @lc app=leetcode.cn id=2437 lang=golang
 *
 * [2437] 有效时间的数目
 */

// @lc code=start
func countTime(time string) int {
	if !strings.Contains(time, "?") {
		return 1
	}

	timeSlice := strings.SplitN(time, ":", 2)
	hour, minute := timeSlice[0], timeSlice[1]
	hourChoice, minuteChoice := 1, 1

	switch strings.Count(hour, "?") {
	case 2:
		hourChoice = 24
	case 1:
		if strings.HasPrefix(hour, "?") {
			if hour[1]-'0' <= 3 {
				hourChoice = 3
			} else {
				hourChoice = 2
			}
		} else {
			if hour[0]-'0' <= 1 {
				hourChoice = 10
			} else {
				hourChoice = 4
			}
		}
	}

	switch strings.Count(minute, "?") {
	case 2:
		minuteChoice = 60
	case 1:
		if strings.HasPrefix(minute, "?") {
			minuteChoice = 6
		} else {
			minuteChoice = 10
		}
	}

	return hourChoice * minuteChoice
}

// @lc code=end

