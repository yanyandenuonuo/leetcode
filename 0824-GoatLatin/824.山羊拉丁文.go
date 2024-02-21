/*
 * @lc app=leetcode.cn id=824 lang=golang
 *
 * [824] 山羊拉丁文
 */

// @lc code=start
func toGoatLatin(sentence string) string {
	resList := strings.Split(sentence, " ")

	for idx, word := range resList {
		firstChar := strings.ToLower(string(word[0]))

		latinWord := &strings.Builder{}
		if firstChar == "a" || firstChar == "e" || firstChar == "i" || firstChar == "o" || firstChar == "u" {
			latinWord.WriteString(word)
		} else {
			latinWord.WriteString(word[1:])
			latinWord.WriteByte(word[0])
		}

		latinWord.WriteString("ma")

		for appendIdx := 0; appendIdx < idx+1; appendIdx += 1 {
			latinWord.WriteRune('a')
		}

		resList[idx] = latinWord.String()
	}

	return strings.Join(resList, " ")
}

// @lc code=end

