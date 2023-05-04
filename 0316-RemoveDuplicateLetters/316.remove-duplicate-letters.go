/*
 * @lc app=leetcode id=316 lang=golang
 *
 * [316] Remove Duplicate Letters
 *
 * https://leetcode.com/problems/remove-duplicate-letters/description/
 *
 * algorithms
 * Medium (38.18%)
 * Likes:    1808
 * Dislikes: 135
 * Total Accepted:    101.7K
 * Total Submissions: 265.8K
 * Testcase Example:  '"bcabc"'
 *
 * Given a string s, remove duplicate letters so that every letter appears once
 * and only once. You must make sure your result is the smallest in
 * lexicographical order among all possible results.
 *
 * Note: This question is the same as 1081:
 * https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
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
 * 1 <= s.length <= 10^4
 * s consists of lowercase English letters.
 *
 *
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
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

