/*
 * @lc app=leetcode.cn id=2446 lang=golang
 *
 * [2446] 判断两个事件是否存在冲突
 */

// @lc code=start
func haveConflict(event1 []string, event2 []string) bool {
	eventMinute1 := []int{
		convertTimeStrToMinute(event1[0]),
		convertTimeStrToMinute(event1[1]),
	}

	eventMinute2 := []int{
		convertTimeStrToMinute(event2[0]),
		convertTimeStrToMinute(event2[1]),
	}

	if eventMinute1[0] <= eventMinute2[0] && eventMinute1[1] >= eventMinute2[0] {
		// 	| 		|
		// 		|		|

		// |			|
		//  	|	|
		return true
	} else if eventMinute1[0] <= eventMinute2[1] && eventMinute1[1] >= eventMinute2[1] {
		// 		| 		|
		// |		|
		return true
	} else if eventMinute1[0] >= eventMinute2[0] && eventMinute1[1] <= eventMinute2[1] {
		// 		|	|
		//  |			|
		return true
	}

	return false
}

func convertTimeStrToMinute(timeStr string) int {
	timeSlice := strings.Split(timeStr, ":")
	hourNum, _ := strconv.Atoi(timeSlice[0])
	minuteNum, _ := strconv.Atoi(timeSlice[1])
	return hourNum*60 + minuteNum
}

// @lc code=end

