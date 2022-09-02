/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (35.12%)
 * Likes:    5258
 * Dislikes: 352
 * Total Accepted:    442.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given a string S and a string T, find the minimum window in S which will
 * contain all the characters in T in complexity O(n).
 * 
 * Example:
 * 
 * 
 * Input: S = "ADOBECODEBANC", T = "ABC"
 * Output: "BANC"
 * 
 * 
 * Note:
 * 
 * 
 * If there is no such window in S that covers all characters in T, return the
 * empty string "".
 * If there is such window, you are guaranteed that there will always be only
 * one unique minimum window in S.
 * 
 * 
 */

// @lc code=start
func minWindow(s string, t string) string {
	targetMap := make(map[string]int, len(t))
	countMap := make(map[string]int, len(t))

	for idx := 0; idx < len(t); idx += 1 {
		currentStr := t[idx:idx+1]
		if _, isExist := targetMap[currentStr]; !isExist {
			targetMap[currentStr]  = 0 
		}
		targetMap[currentStr] += 1
	}

	res := ""
	isMatchCount := 0

	for left, right := 0, 0; right < len(s); right += 1 {
		currentStr := s[right:right+1]

		if _, isExist := targetMap[currentStr]; !isExist {
			continue
		}

		if _, isExist := countMap[currentStr]; !isExist {
			countMap[currentStr] = 0
		}
		countMap[currentStr] += 1

		if countMap[currentStr] == targetMap[currentStr] {
			isMatchCount += 1
		}


		for isMatchCount == len(targetMap) {
			if right - left + 1 < len(res) || len(res) == 0 {
				res = s[left:right+1]
			}
			currentStr := s[left:left+1]
			if _, isExist := targetMap[currentStr]; isExist {
				countMap[currentStr] -= 1

				if countMap[currentStr] < targetMap[currentStr] {
					isMatchCount -= 1
				}
			}
			
			left += 1
		}
	}

	return res
}

// @lc code=end

