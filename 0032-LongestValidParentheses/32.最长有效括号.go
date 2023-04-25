/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	// solution1: 动态规划
	/**
	if len(s) < 2 {
		return 0
	}

	maxLength := 0
	dp := make([]int, len(s))

	for idx := 0; idx < len(s); idx += 1 {
		if s[idx] == '(' {
			continue
		}
		if idx-1 < 0 {
			continue
		}

		if s[idx-1] == '(' {
			if idx-2 >= 0 {
				dp[idx] = dp[idx-2] + 2
			} else {
				dp[idx] = 2
			}
		} else if s[idx-1] == ')' && idx-dp[idx-1]-1 >= 0 && s[idx-dp[idx-1]-1] == '(' {
			if idx-dp[idx-1]-2 >= 0 {
				dp[idx] = dp[idx-dp[idx-1]-2] + dp[idx-1] + 2
			} else {
				dp[idx] = dp[idx-1] + 2
			}
		}
		if dp[idx] > maxLength {
			maxLength = dp[idx]
		}
	}

	return maxLength
	*/

	// solution2: 使用栈
	//			  ( 直接入栈
	//			  ) 若栈为空则入栈
	//				若栈不为空，尝试与栈顶元素匹配
	//					若不匹配则直接入栈
	//					若匹配则弹出栈顶元素，当前idx - 当前栈顶元素idx即为当前的匹配长度
	/**
	stackIdx := make([]int, 0, len(s))
	// 消除栈为空的场景
	stackIdx = append(stackIdx, -1)

	maxLength := 0
	for idx := 0; idx < len(s); idx += 1 {
		if s[idx] == '(' {
			stackIdx = append(stackIdx, idx)
			continue
		}

		// 实质指前面字符均已完成匹配
		if stackIdx[len(stackIdx)-1] < 0 {
			stackIdx = append(stackIdx, idx)
			continue
		}

		if s[stackIdx[len(stackIdx)-1]] == ')' {
			stackIdx = append(stackIdx, idx)
			continue
		}

		stackIdx = stackIdx[:len(stackIdx)-1]
		currentLength := idx - stackIdx[len(stackIdx)-1]
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
	*/

	// solution3: 贪心，利用leftCount和rightCount计数
	//			  从前往后遍历，若rightCount > leftCount则计数清零
	//						 若rightCount == leftCount则leftCount + rightCount即为匹配长度
	//			  针对(()这种case则需要从后往前遍历
	//			  从后往前遍历，若leftCount > rightCount则计数清零
	//						 若rightCount == leftCount则leftCount + rightCount即为匹配长度
	leftCount, rightCount := 0, 0
	maxLength := 0
	for idx := 0; idx < len(s); idx += 1 {
		switch s[idx] {
		case '(':
			leftCount += 1
		case ')':
			rightCount += 1
		}
		if rightCount > leftCount {
			leftCount = 0
			rightCount = 0
			continue
		}

		if leftCount == rightCount && leftCount+rightCount > maxLength {
			maxLength = leftCount + rightCount
		}
	}

	// 针对(()的case则尝试从右边开始遍历
	leftCount, rightCount = 0, 0
	for idx := len(s) - 1; idx >= 0; idx -= 1 {
		switch s[idx] {
		case '(':
			leftCount += 1
		case ')':
			rightCount += 1
		}

		if leftCount == rightCount && leftCount+rightCount > maxLength {
			maxLength = leftCount + rightCount
		}
		if leftCount > rightCount {
			leftCount = 0
			rightCount = 0
		}
	}

	return maxLength
}

// @lc code=end

