/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	for leftIdx, rightIdx := 0, len(s)-1; leftIdx < rightIdx; {
		leftChar, isValidChar := getRuneChar(s[leftIdx])
		if !isValidChar {
			leftIdx += 1
			continue
		}

		rightChar, isValidChar := getRuneChar(s[rightIdx])
		if !isValidChar {
			rightIdx -= 1
			continue
		}

		if leftChar != rightChar {
			return false
		}
		leftIdx += 1
		rightIdx -= 1
	}
	return true
}

func getRuneChar(runeChar byte) (byte, bool) {
	if runeChar >= 'A' && runeChar <= 'Z' {
		runeChar = runeChar - 'A' + 'a'
	}
	if (runeChar >= 'a' && runeChar <= 'z') || (runeChar >= '0' && runeChar <= '9') {
		return runeChar, true
	}
	return ' ', false
}

// @lc code=end

