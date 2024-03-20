/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(s string) string {
	// solution: 模拟栈
	stack := make([]string, 0, len(s))
	for idx := range s {
		skipCurrent := false

		for len(stack) > 0 && stack[len(stack)-1] == s[idx:idx+1] {
			skipCurrent = true
			stack = stack[:len(stack)-1]
		}

		if !skipCurrent {
			stack = append(stack, s[idx:idx+1])
		}
	}

	return strings.Join(stack, "")
}

// @lc code=end

