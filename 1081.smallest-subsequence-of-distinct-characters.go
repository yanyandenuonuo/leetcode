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

