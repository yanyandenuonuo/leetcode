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
	countMap := make(map[string]int, len(s))
	handleMap := make(map[string]bool, len(s))
    for idx := 0; idx < len(s); idx += 1 {
		countMap[s[idx:idx+1]] += 1
	}

	res := ""
    for idx := 0; idx < len(s); idx += 1 {
		val := s[idx:idx+1]
		countMap[val] -= 1
		if _, isHandle := handleMap[val]; isHandle {
			continue
		}

		for len(res) > 0 && countMap[res[len(res)-1:]] > 0 && res[len(res)-1:] > val {
			delete(handleMap, res[len(res)-1:])
			res = res[:len(res)-1]
		}

		res += val
		handleMap[val] = true
	}
	return res
}
// @lc code=end

