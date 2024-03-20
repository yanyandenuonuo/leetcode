/*
 * @lc app=leetcode.cn id=1078 lang=golang
 *
 * [1078] Bigram 分词
 */

// @lc code=start
func findOcurrences(text string, first string, second string) []string {
	// solution1: 先分词，再匹配两个单词
	// solution2: 合并单词为一个，按字符串匹配后取后一个单词
	matchStr := fmt.Sprintf(" %s %s ", first, second)
	text = fmt.Sprintf(" %s", text) // 强制first前也应该有一个空格

	resList := make([]string, 0)

	for idx := 0; idx < len(text)-len(matchStr); idx += 1 {
		if text[idx:idx+len(matchStr)] != matchStr {
			continue
		}

		endIdx := idx + len(matchStr)

		for endIdx < len(text) && text[endIdx] != ' ' {
			endIdx += 1
		}

		resList = append(resList, text[idx+len(matchStr):endIdx])
	}

	return resList
}

// @lc code=end

