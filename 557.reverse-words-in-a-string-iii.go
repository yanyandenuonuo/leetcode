/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func reverseWords(s string) string {
	stringRune := []rune(s)
	subLeftIdx := 0
	for idx := 0; idx < len(stringRune); idx += 1 {
		if stringRune[idx] != ' ' && idx < (len(stringRune) - 1) {
			continue
		}
		
		subRightIdx := idx - 1
		if idx == (len(stringRune) - 1) {
			subRightIdx = idx
		}
		
		for subLeftIdx < subRightIdx {
			stringRune[subLeftIdx], stringRune[subRightIdx] = stringRune[subRightIdx], stringRune[subLeftIdx]
			subLeftIdx += 1
			subRightIdx -= 1
		}
		subLeftIdx = idx + 1
	}
	return string(stringRune)
}
// @lc code=end

