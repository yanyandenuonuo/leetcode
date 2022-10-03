/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	mockStack := make([]string, 0, len(s))
	for idx := 0; idx < len(s); idx += 1 {
		currentChar := s[idx : idx+1]
		switch currentChar {
		case "(":
			fallthrough
		case "{":
			fallthrough
		case "[":
			mockStack = append(mockStack, currentChar)
		case ")":
			if len(mockStack) > 0 && mockStack[len(mockStack)-1] == "(" {
				mockStack = mockStack[:len(mockStack)-1]
			} else {
				return false
			}
		case "}":
			if len(mockStack) > 0 && mockStack[len(mockStack)-1] == "{" {
				mockStack = mockStack[:len(mockStack)-1]
			} else {
				return false
			}
		case "]":
			if len(mockStack) > 0 && mockStack[len(mockStack)-1] == "[" {
				mockStack = mockStack[:len(mockStack)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}
	if len(mockStack) == 0 {
		return true
	}

	return false
}

// @lc code=end

