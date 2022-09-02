/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte)  {
    rightIdx := len(s) - 1
	for leftIdx := 0; leftIdx < rightIdx; {
		s[leftIdx], s[rightIdx] = s[rightIdx], s[leftIdx]
		leftIdx += 1
		rightIdx -= 1
	}
}
// @lc code=end

