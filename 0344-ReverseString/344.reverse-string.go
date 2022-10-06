/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte) {
	for leftIdx, rightIdx := 0, len(s)-1; leftIdx < rightIdx; leftIdx += 1 {
		s[leftIdx], s[rightIdx] = s[rightIdx], s[leftIdx]
		rightIdx -= 1
	}
}

// @lc code=end

