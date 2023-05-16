/*
 * @lc app=leetcode.cn id=424 lang=golang
 *
 * [424] 替换后的最长重复字符
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	// solution: 双指针滑动窗口
	maxLength, maxChar, charCounter := 0, byte(0), make(map[byte]int)
	for leftIdx, rightIdx := 0, 0; rightIdx < len(s); rightIdx += 1 {
		charCounter[s[rightIdx]] += 1

		if charCounter[s[rightIdx]] > charCounter[maxChar] {
			maxChar = s[rightIdx]
		}

		if rightIdx-leftIdx+1-charCounter[maxChar] <= k {
			if rightIdx-leftIdx+1 > maxLength {
				maxLength = rightIdx - leftIdx + 1
			}

			continue
		}

		charCounter[s[leftIdx]] -= 1
		leftIdx += 1
	}

	return maxLength
}

// @lc code=end

