/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitMap := map[uint8][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	res := make([]string, 0, 4*4*3*3)

	var backtrace func(int, string)
	backtrace = func(idx int, combination string) {
		if idx == len(digits) {
			res = append(res, combination)
			return
		}

		for _, choiceChar := range digitMap[digits[idx]] {
			backtrace(idx+1, combination+choiceChar)
		}
	}

	backtrace(0, "")
	return res
}

// @lc code=end

