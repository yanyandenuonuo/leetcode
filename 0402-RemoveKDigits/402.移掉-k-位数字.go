/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉 K 位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	// solution: 贪心+单调栈
	stack := make([]rune, 0)

	for _, runeChar := range num {
		for len(stack) > 0 && runeChar < stack[len(stack)-1] && k > 0 {
			stack = stack[:len(stack)-1]
			k -= 1
		}

		stack = append(stack, runeChar)
	}

	// 移除最大的k个数，移除前导0
	startIdx := 0
	endIdx := len(stack) - k

	for ; startIdx < endIdx && stack[startIdx] == '0'; startIdx += 1 {
	}

	if startIdx >= endIdx {
		return "0"
	}

	return string(stack[startIdx:endIdx])
}

// @lc code=end

