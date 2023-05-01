/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	for idx, lowerCounter, upperCounter := 0, 0, 0; idx < len(word); idx += 1 {
		if idx == 0 && word[idx] >= 'A' && word[idx] <= 'Z' {
			continue
		}
		if word[idx] >= 'A' && word[idx] <= 'Z' {
			upperCounter += 1
		} else {
			lowerCounter += 1
		}

		if lowerCounter > 0 && upperCounter > 0 {
			return false
		}
	}

	return true
}

// @lc code=end

