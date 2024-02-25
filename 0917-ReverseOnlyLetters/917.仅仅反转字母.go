/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
func reverseOnlyLetters(s string) string {
	stringList := make([]string, len(s))
	for leftIdx, rightIdx := 0, len(s)-1; leftIdx <= rightIdx; {
		if s[leftIdx] < 'A' || ('Z' < s[leftIdx] && s[leftIdx] < 'a') || 'z' < s[leftIdx] {
			stringList[leftIdx] = s[leftIdx : leftIdx+1]
			leftIdx += 1

			continue
		}

		if s[rightIdx] < 'A' || ('Z' < s[rightIdx] && s[rightIdx] < 'a') || 'z' < s[rightIdx] {
			stringList[rightIdx] = s[rightIdx : rightIdx+1]

			rightIdx -= 1
			continue
		}

		stringList[leftIdx] = s[rightIdx : rightIdx+1]
		stringList[rightIdx] = s[leftIdx : leftIdx+1]

		leftIdx += 1
		rightIdx -= 1
	}

	return strings.Join(stringList, "")
}

// @lc code=end

