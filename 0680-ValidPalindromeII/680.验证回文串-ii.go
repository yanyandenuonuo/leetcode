/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文串 II
 */

// @lc code=start
func validPalindrome(s string) bool {
	// solution: 双指针
	for leftIdx, rightIdx := 0, len(s)-1; leftIdx < rightIdx; {
		if s[leftIdx] == s[rightIdx] {
			leftIdx += 1
			rightIdx -= 1
			continue
		}

		// 尝试删左边，注意不要改变原有leftIdx和rightIdx，不然会影响到尝试删右边
		isPalindrome := true
		for newLeftIdx, newRightIdx := leftIdx+1, rightIdx; newLeftIdx < newRightIdx; {
			if s[newLeftIdx] != s[newRightIdx] {
				isPalindrome = false
				break
			}

			newLeftIdx += 1
			newRightIdx -= 1
		}

		if isPalindrome {
			break
		}

		// 尝试删右边，注意reset isPalindrome
		isPalindrome = true
		for newLeftIdx, newRightIdx := leftIdx, rightIdx-1; newLeftIdx < newRightIdx; {
			if s[newLeftIdx] != s[newRightIdx] {
				isPalindrome = false
				break
			}

			newLeftIdx += 1
			newRightIdx -= 1
		}

		if isPalindrome {
			break
		} else {
			return false
		}
	}

	return true
}

// @lc code=end

