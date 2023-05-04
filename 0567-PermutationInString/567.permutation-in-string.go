/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 *
 * https://leetcode.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (44.44%)
 * Likes:    1860
 * Dislikes: 70
 * Total Accepted:    153.1K
 * Total Submissions: 344.5K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * Given two strings s1 and s2, write a function to return true if s2 contains
 * the permutation of s1. In other words, one of the first string's
 * permutations is the substring of the second string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s1 = "ab" s2 = "eidbaooo"
 * Output: True
 * Explanation: s2 contains one permutation of s1 ("ba").
 *
 *
 * Example 2:
 *
 *
 * Input:s1= "ab" s2 = "eidboaoo"
 * Output: False
 *
 *
 *
 * Constraints:
 *
 *
 * The input strings only contain lower case letters.
 * The length of both given strings is in range [1, 10,000].
 *
 *
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	// solution1: 双指针，查看匹配字符的长度是否刚好为len(s1)
	/**
	targetMap := make(map[string]int, len(s1))
	countMap := make(map[string]int, len(s1))

	for idx := 0; idx < len(s1); idx += 1 {
		currentStr := s1[idx : idx+1]
		if _, isExist := targetMap[currentStr]; !isExist {
			targetMap[currentStr] = 0
		}
		targetMap[currentStr] += 1
	}

	isMatchCount := 0

	for left, right := 0, 0; right < len(s2); right += 1 {
		currentStr := s2[right : right+1]

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
			if len(s2[left:right+1]) == len(s1) {
				return true
			}
			currentStr := s2[left : left+1]
			if _, isExist := targetMap[currentStr]; isExist {
				countMap[currentStr] -= 1

				if countMap[currentStr] < targetMap[currentStr] {
					isMatchCount -= 1
				}
			}

			left += 1
		}
	}

	return false
	*/

	// solution2: 滑动窗口，窗口长度len(s1)的字符是否匹配
	s1Idx := [26]int{}
	for _, runeChar := range s1 {
		s1Idx[runeChar-'a'] += 1
	}

	s2Idx := [26]int{}
	for idx := 0; idx < len(s2); idx += 1 {
		if idx < len(s1) {
			s2Idx[s2[idx]-'a'] += 1
			continue
		}

		if s1Idx == s2Idx {
			return true
		}

		s2Idx[s2[idx-len(s1)]-'a'] -= 1
		s2Idx[s2[idx]-'a'] += 1
	}

	return s1Idx == s2Idx
}

// @lc code=end

