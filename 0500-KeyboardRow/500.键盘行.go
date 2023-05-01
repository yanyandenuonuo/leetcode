/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
func findWords(words []string) []string {
	// solution: 对字符做与行的映射，然后逐个单词比对
	//			 字符映射可以做map，也可用数组或字符串
	//			 'a' => 1, b => '2', ..., z => 2
	//			 []int{1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2}
	//			 12210111011122000010020202
	runeIdx := [26]int{1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2}
	res := make([]string, 0, len(words))
	for idx, word := range words {
		if len(word) == 0 {
			continue
		}

		word = strings.ToLower(word)
		currentIdx := runeIdx[word[0]-'a']
		charIdx := 1
		for ; charIdx < len(word); charIdx += 1 {
			if runeIdx[word[charIdx]-'a'] != currentIdx {
				break
			}
		}

		if charIdx == len(word) {
			res = append(res, words[idx])
		}
	}
	return res
}

// @lc code=end

