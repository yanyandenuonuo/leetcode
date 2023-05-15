/*
 * @lc app=leetcode.cn id=709 lang=golang
 *
 * [709] 转换成小写字母
 */

// @lc code=start
func toLowerCase(s string) string {
	// solution: 逐个字符转换
	//			 大写字符范围为[65, 90]，小写字符范围为[97, 122]，大写字符+32即为小写字符，观察实际二进制表示实际+32可转换为|32
	strRunes := make([]rune, len(s))
	for idx, charRune := range s {
		if charRune >= 'A' && charRune <= 'Z' {
			// charRune = charRune - 'A' + 'a'
			// charRune += 32
			charRune |= 32
		}

		strRunes[idx] = charRune
	}

	return string(strRunes)
}

// @lc code=end

