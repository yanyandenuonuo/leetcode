/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// solution1: 横向扫描，比较两个字符求最长公共前缀，再依次与后续字符比较
	// solution2: 按公共字符前缀逐个字符比较
	res := ""
	for idx := 0; ; idx += 1 {
		if idx >= len(strs[0]) {
			return res
		}
		currentChar := strs[0][idx]

		for _, str := range strs {
			if idx >= len(str) || str[idx] != currentChar {
				return res
			}
		}
		res += string(currentChar)
	}
	return res
}

// @lc code=end

