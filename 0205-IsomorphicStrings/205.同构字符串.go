/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sourceMap := make(map[string]string, len(s))
	targetMap := make(map[string]string, len(s))
	for idx := 0; idx < len(s); idx += 1 {
		curChar := s[idx : idx+1]
		if _, isExist := sourceMap[curChar]; !isExist {
			if _, isExist := targetMap[t[idx:idx+1]]; isExist {
				return false
			}
			sourceMap[curChar] = t[idx : idx+1]
			targetMap[t[idx:idx+1]] = curChar
		} else if sourceMap[curChar] != t[idx:idx+1] {
			return false
		}
	}

	return true
}

// @lc code=end

