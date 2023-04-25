/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */

// @lc code=start
func readBinaryWatch(turnedOn int) []string {
	// solution1: 枚举时分数据，表示时分的LED和等于turnedOn的即为答案
	res := make([]string, 0)
	for hourNum := 0; hourNum < 12; hourNum += 1 {
		bitHour := bitCount(hourNum)
		if bitHour > turnedOn {
			continue
		}
		for minuteNum := 0; minuteNum < 60; minuteNum += 1 {
			if bitHour+bitCount(minuteNum) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", hourNum, minuteNum))
			}
		}
	}
	return res

	// solution2: 以二进制方式枚举，二进制中1的位数等于turnedOn的即为可选答案
	/**
	res := make([]string, 0)
	for idx := 0; idx < 1<<10; idx += 1 {
		minuteNum := idx & (1<<6 - 1)
		if minuteNum > 59 {
			continue
		}

		hourNum := idx >> 6
		if hourNum > 11 {
			continue
		}
		if bitCount(minuteNum)+bitCount(hourNum) == turnedOn {
			res = append(res, fmt.Sprintf("%d:%02d", hourNum, minuteNum))
		}

	}
	return res
	*/
}

func bitCount(num int) int {
	count := 0
	for ; num > 0; num >>= 1 {
		if num&0x01 == 0x01 {
			count += 1
		}
	}
	return count
}

// @lc code=end

