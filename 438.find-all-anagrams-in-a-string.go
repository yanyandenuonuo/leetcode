/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 *
 * https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (43.92%)
 * Likes:    3511
 * Dislikes: 183
 * Total Accepted:    304.1K
 * Total Submissions: 691.9K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * Given a string s and a non-empty string p, find all the start indices of p's
 * anagrams in s.
 * 
 * Strings consists of lowercase English letters only and the length of both
 * strings s and p will not be larger than 20,100.
 * 
 * The order of output does not matter.
 * 
 * Example 1:
 * 
 * Input:
 * s: "cbaebabacd" p: "abc"
 * 
 * Output:
 * [0, 6]
 * 
 * Explanation:
 * The substring with start index = 0 is "cba", which is an anagram of "abc".
 * The substring with start index = 6 is "bac", which is an anagram of
 * "abc".
 * 
 * 
 * 
 * Example 2:
 * 
 * Input:
 * s: "abab" p: "ab"
 * 
 * Output:
 * [0, 1, 2]
 * 
 * Explanation:
 * The substring with start index = 0 is "ab", which is an anagram of "ab".
 * The substring with start index = 1 is "ba", which is an anagram of "ab".
 * The substring with start index = 2 is "ab", which is an anagram of "ab".
 * 
 * 
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
    targetMap := make(map[string]int, len(p))
	countMap := make(map[string]int, len(p))

	for idx := 0; idx < len(p); idx += 1 {
		currentStr := p[idx:idx+1]
		if _, isExist := targetMap[currentStr]; !isExist {
			targetMap[currentStr]  = 0 
		}
		targetMap[currentStr] += 1
	}

	isMatchCount := 0

	res := make([]int, 0, len(s))

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
			currentMatch := s[left:right+1]
			if len(currentMatch) == len(p) {
				res = append(res, left)
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

