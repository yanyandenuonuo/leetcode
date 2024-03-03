/*
 * @lc app=leetcode.cn id=2451 lang=golang
 *
 * [2451] 差值数组不同的字符串
 */

// @lc code=start
func oddString(words []string) string {
	for idx := 1; idx < len(words[0]); idx += 1 {
		diffMap := make(map[int][]int, 2)
		for wordIdx, word := range words {
			diff := int(word[idx] - word[idx-1])
			if _, isExist := diffMap[diff]; !isExist {
				diffMap[diff] = make([]int, 0, len(words[0])-1)
			}

			diffMap[diff] = append(diffMap[diff], wordIdx)

			if len(diffMap) > 1 && wordIdx > 1 {
				for key := range diffMap {
					if len(diffMap[key]) == 1 {
						return words[diffMap[key][0]]
					}
				}
			}
		}
	}

	return ""
}

// @lc code=end

