/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func reverseWords(s string) string {
	sRunes := []rune(s)
	for idx := 0; idx < len(sRunes); {
		if sRunes[idx] == ' ' {
			idx += 1
			continue
		}

		nextIdx := idx + 1
		for ; nextIdx < len(sRunes) && sRunes[nextIdx] != ' '; nextIdx += 1 {

		}

		for rightIdx := nextIdx - 1; idx < rightIdx; idx += 1 {
			sRunes[idx], sRunes[rightIdx] = sRunes[rightIdx], sRunes[idx]
			rightIdx -= 1
		}

		idx = nextIdx
	}

	return string(sRunes)
}

// @lc code=end

