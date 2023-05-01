/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	// solution1: 枚举重复子串的字符长度
	// solution2: 字符串匹配，将2个s拼接成ss，前后移除一部分字符后，若依然能查找到s则符合条件
	return strings.Index((s + s)[1:2*len(s)-1], s) != -1
	// solution3: kmp算法
	// todo
}

// @lc code=end

