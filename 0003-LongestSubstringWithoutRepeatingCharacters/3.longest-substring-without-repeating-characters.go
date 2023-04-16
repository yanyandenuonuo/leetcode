/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (30.71%)
 * Likes:    11218
 * Dislikes: 625
 * Total Accepted:    1.8M
 * Total Submissions: 5.7M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string s, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * Notice that the answer must be a substring, "pwke" is a subsequence and not
 * a substring.
 *
 *
 * Example 4:
 *
 *
 * Input: s = ""
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s consists of English letters, digits, symbols and spaces.
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	charCountMap := make(map[uint8]int, len(s))

	maxLength := 0

	for left, right := 0, 0; right < len(s); right += 1 {
		charCountMap[s[right]] += 1

		for charCountMap[s[right]] > 1 {
			charCountMap[s[left]] -= 1
			left += 1
		}

		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}

// @lc code=end

