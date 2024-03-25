/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
func reverseWords(s string) string {
	// solution1: 按空格拆单词放数组，逆序拼凑
	wordList := make([]string, 0)
	for idx := 0; idx < len(s); {
		if s[idx] == ' ' {
			idx += 1
			continue
		}

		rightIdx := idx + 1
		for rightIdx < len(s) && s[rightIdx] != ' ' {
			rightIdx += 1
		}

		wordList = append(wordList, s[idx:rightIdx])

		idx = rightIdx + 1
	}

	stringBuilder := new(strings.Builder)

	for idx := len(wordList) - 1; idx >= 0; idx -= 1 {
		stringBuilder.WriteString(wordList[idx])
		if idx != 0 {
			stringBuilder.WriteRune(' ')
		}
	}

	return stringBuilder.String()

	// solution2: 先全部翻转，再按单词逐个翻转
	// solution3: 按单词翻转，双指针记录位置，使用额外空间存放较长的单词
}

// @lc code=end

