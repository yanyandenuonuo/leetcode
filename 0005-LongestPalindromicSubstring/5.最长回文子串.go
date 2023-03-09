/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	// solution1
	// 动态规划
	// dp[i][j]表示s[i:j+1]是否为回文字符串
	// 1. dp[i][i] = true
	// 2. s[i] != s[j]
	//		dp[i][j] = false
	// 3. s[i] == s[j]
	//		j-i < 2 	=> dp[i][j] = true
	// 		j-i >= 2 	=> dp[i][j] = dp[i+1][j-1]
	// 因为长子串依赖短子串，所以需要基于字串长度作为一个基础迭代

	// if len(s) < 2 {
	// 	return s
	// }
	// dp := make([][]bool, len(s))
	// for idx := 0; idx < len(dp); idx += 1 {
	// 	dp[idx] = make([]bool, len(dp))
	// 	dp[idx][idx] = true
	// }

	// maxLen := 1
	// startIdx := 0
	// for offset := 1; offset < len(s); offset += 1 {
	// 	for leftIdx := 0; leftIdx < len(s); leftIdx += 1 {
	// 		rightIdx := leftIdx + offset
	// 		if rightIdx == len(s) {
	// 			break
	// 		}

	// 		if s[leftIdx] != s[rightIdx] {
	// 			dp[leftIdx][rightIdx] = false
	// 			continue
	// 		}
	// 		if rightIdx-leftIdx < 2 {
	// 			dp[leftIdx][rightIdx] = true
	// 		} else {
	// 			dp[leftIdx][rightIdx] = dp[leftIdx+1][rightIdx-1]
	// 		}
	// 		if dp[leftIdx][rightIdx] && (rightIdx-leftIdx+1) > maxLen {
	// 			maxLen = rightIdx - leftIdx + 1
	// 			startIdx = leftIdx
	// 		}
	// 	}
	// }
	// return s[startIdx : startIdx+maxLen]

	// solution2
	// 中心扩展法
	// 选定一个点作为回文子串中心向两边扩展，求最大，需要考虑中心点为奇数和偶数的情形
	// if len(s) < 2 {
	// 	return s
	// }

	// maxLen := 1
	// startIdx := 0
	// for checkIdx := 0; checkIdx < len(s); checkIdx += 1 {
	// 	oddLength := checkLongestPalindrome(s, checkIdx, checkIdx)
	// 	evenLength := checkLongestPalindrome(s, checkIdx, checkIdx+1)

	// 	if oddLength > maxLen {
	// 		maxLen = oddLength
	// 		startIdx = checkIdx - (maxLen-1)/2
	// 	}
	// 	if evenLength > maxLen {
	// 		maxLen = evenLength
	// 		startIdx = checkIdx - (maxLen-1)/2
	// 	}
	// }
	// return s[startIdx : startIdx+maxLen]

	// solution3
	// Manacher算法，利用空间换时间，基于最大回文字符串对称性来进行时间优化
	// 先利用插值消除奇偶性问题，aba => #a#b#a#，abba => #a#b#b#a#
	// scanIdx > maxRight, 利用中心扩展法求解
	// scanIdx < maxRight, mirrorIdx = 2*centerIdx-scanIdx, rMap[scanIdx] = min(rMap[mirrorIdx], maxRight-idx)，然后再进行扩展
	// 		mirrorIdx = centerIdx - (scanIdx - centerIdx) = 2*centerIdx-scanIdx
	if len(s) < 2 {
		return s
	}

	rMap := make(map[int]int, len(s))
	sBytes := make([]rune, 0, 2*len(s)+1)
	sBytes = append(sBytes, '#')
	for _, scanRune := range s {
		sBytes = append(sBytes, scanRune, '#')
	}
	sWithSeparator := string(sBytes)

	maxCenter, maxRight := 1, -1
	for idx := 0; idx < len(sWithSeparator); idx++ {
		if idx > maxRight {
			rMap[idx] = (checkLongestPalindrome(sWithSeparator, idx, idx) - 1) / 2
		} else {
			mirrorIdx := 2*maxCenter - idx
			currentR := rMap[mirrorIdx]
			if currentR > maxRight-idx {
				currentR = maxRight - idx
			}
			rMap[idx] = (checkLongestPalindrome(sWithSeparator, idx-currentR, idx+currentR) - 1) / 2
		}
		if rMap[idx] > rMap[maxCenter] {
			maxCenter = idx
			maxRight = idx + rMap[maxCenter]
		}
	}
	resBytes := make([]rune, 0, rMap[maxCenter]/2)
	for idx, scanRune := range sWithSeparator[maxCenter-rMap[maxCenter] : maxCenter+rMap[maxCenter]+1] {
		if idx&0x01 == 0x01 {
			resBytes = append(resBytes, scanRune)
		}
	}

	return string(resBytes)
}

func checkLongestPalindrome(s string, leftIdx, rightIdx int) int {
	maxLen := 0
	for leftIdx >= 0 && rightIdx < len(s) && s[leftIdx] == s[rightIdx] {
		maxLen = rightIdx - leftIdx + 1
		leftIdx -= 1
		rightIdx += 1
	}
	return maxLen
}

// @lc code=end

