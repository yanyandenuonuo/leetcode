/*
 * @lc app=leetcode.cn id=1576 lang=golang
 *
 * [1576] 替换所有的问号
 */

// @lc code=start
func modifyString(s string) string {
	// solution: 逐个替换，注意偏移量溢出
	stringList := make([]string, len(s))
	for idx := range s {
		if s[idx] != '?' {
			stringList[idx] = s[idx : idx+1]
			continue
		}

		if idx == 0 {
			if idx+1 < len(s) {
				stringList[idx] = string('a' + rune(s[idx+1]+25)%26)
			} else {
				stringList[idx] = "a"
			}
		} else if idx == len(s)-1 {
			stringList[idx] = string('a' + (stringList[idx-1][0]+1)%26)
		} else {
			offset := byte(1)

			for 'a'+(stringList[idx-1][0]+offset)%26 == s[idx+1] {
				offset += 1
			}
			stringList[idx] = string('a' + (stringList[idx-1][0]+offset)%26)
		}
	}

	return strings.Join(stringList, "")
}

// @lc code=end

