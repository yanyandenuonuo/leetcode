/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(s string, goal string) bool {
	// solution: s与goal完全相等且s至少存在一组相同字符对或者s与goal只有2个字符不等且交换后相等

	if len(s) != len(goal) {
		return false
	}

	charCount := make([]rune, 26)
	diffIdxList := make([]int, 0, len(s))

	for idx, _ := range s {
		charCount[s[idx]-'a'] += 1

		if s[idx] == goal[idx] {
			continue
		}

		diffIdxList = append(diffIdxList, idx)

		if len(diffIdxList) > 2 {
			return false
		}
	}

	switch len(diffIdxList) {
	case 0:
		for _, repeatCount := range charCount {
			if repeatCount > 1 {
				return true
			}
		}
		return false
	case 2:
		return s[diffIdxList[0]] == goal[diffIdxList[1]] && s[diffIdxList[1]] == goal[diffIdxList[0]]
	default:
		return false
	}
}

// @lc code=end

