/*
 * @lc app=leetcode.cn id=2452 lang=golang
 *
 * [2452] 距离字典两次编辑以内的单词
 */

// @lc code=start
func twoEditWords(queries []string, dictionary []string) []string {
	// solution: 枚举
	res := make([]string, 0, len(queries))
	for _, word := range queries {
		for _, dict := range dictionary {
			diffCount := 0
			for idx := 0; idx < len(word); idx += 1 {
				if int(dict[idx]-word[idx]) != 0 {
					diffCount += 1
				}

				if diffCount > 2 {
					break
				}
			}

			if diffCount <= 2 {
				res = append(res, word)
				break
			}
		}
	}
	return res
}

// @lc code=end

