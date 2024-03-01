/*
 * @lc app=leetcode.cn id=1185 lang=golang
 *
 * [1185] 一周中的第几天
 */

// @lc code=start
func dayOfTheWeek(day int, month int, year int) string {
	// solution1: 找参照物，1970年12月31日是星期四
	weekStrList := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	days := (year-1971)*365 + (year-1969)/4

	for monthIdx := 1; monthIdx < month; monthIdx += 1 {
		switch monthIdx {
		case 1, 3, 5, 7, 8, 10, 12:
			days += 31
		case 2:
			days += 28
		default:
			days += 30
		}
	}

	// 当年为闰年的影响
	if month >= 3 && (year%400 == 0 || (year%4 == 0 && year%100 != 0)) {
		days += 1
	}

	return weekStrList[(days+day+4)%7]

	// solution2: 基姆拉尔森计算公式，m的取值范围为3至14（某年的1、2月要看作上一年的13、14月来计算）
	//			  W=(d+2*m+3*(m+1)/5+y+y/4-y/100+y/400+1)%7

	/**
	weekStrList := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	if month < 3 {
		month += 12
		year -= 1
	}

	return weekStrList[(day+2*month+3*(month+1)/5+year+year/4-year/100+year/400+1)%7]
	*/

	// solution3: 蔡勒公式，c年份前两位数 y年份后两位数，m的取值范围为3至14（某年的1、2月要看作上一年的13、14月来计算）
	//			  1582.10.4及之前 w=(y+y/4+c/4-2*c+2m+(3*(m+1))/5+d+1)%7
	//			  1582.10.4之后 w=(y+y/4+c/4-2*c+(26*(m+1))/10+d-1)%7
}

// @lc code=end

