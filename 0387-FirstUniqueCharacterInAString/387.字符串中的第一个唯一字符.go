/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	// solution1: 使用hash表存储词频
	// solution2: 使用hash表存储索引，多次出现将索引设置为-1
	charIdxMap := make(map[rune]int, len(s))
	for idx, runeChar := range s {
		if _, isExist := charIdxMap[runeChar]; !isExist {
			charIdxMap[runeChar] = idx
			continue
		}
		charIdxMap[runeChar] = -1
	}
	for idx, runeChar := range s {
		if charIdxMap[runeChar] != -1 {
			return idx
		}
	}
	return -1
}

// @lc code=end

