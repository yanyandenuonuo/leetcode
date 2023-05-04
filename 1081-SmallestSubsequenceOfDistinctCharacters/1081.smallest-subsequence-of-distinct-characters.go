/*
 * @lc app=leetcode id=1081 lang=golang
 *
 * [1081] Smallest Subsequence of Distinct Characters
 *
 * https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/description/
 *
 * algorithms
 * Medium (53.00%)
 * Likes:    632
 * Dislikes: 92
 * Total Accepted:    15.8K
 * Total Submissions: 29.7K
 * Testcase Example:  '"bcabc"'
 *
 * Return the lexicographically smallest subsequence of s that contains all the
 * distinct characters of s exactly once.
 *
 * Note: This question is the same as 316:
 * https://leetcode.com/problems/remove-duplicate-letters/
 *
 *
 * Example 1:
 *
 *
 * Input: s = "bcabc"
 * Output: "abc"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "cbacdcbc"
 * Output: "acdb"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s consists of lowercase English letters.
 *
 *
 */

// @lc code=start
func smallestSubsequence(s string) string {
	// solution: 贪心+单调栈
	charCounter := [26]int{}
	for _, runeChar := range s {
		charCounter[runeChar-'a'] += 1
	}

	charStack := make([]rune, 0, len(s))
	stackMem := [26]bool{}

	for _, runeChar := range s {
		// 已经在栈里就不再处理，直接更新字符计数即可
		if !stackMem[runeChar-'a'] {
			// 移除stack中存在重复且字典序靠后的字符
			for len(charStack) > 0 && runeChar < charStack[len(charStack)-1] &&
				charCounter[charStack[len(charStack)-1]-'a'] > 0 {
				stackMem[charStack[len(charStack)-1]-'a'] = false
				charStack = charStack[:len(charStack)-1]
			}

			charStack = append(charStack, runeChar)
			stackMem[runeChar-'a'] = true
		}

		charCounter[runeChar-'a'] -= 1
	}

	return string(charStack)
}

// @lc code=end

