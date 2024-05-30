/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	// solution: 模拟+双指针
	stack := make([]int, 0, len(pushed))
	for pushIdx, popIdx := 0, 0; pushIdx < len(pushed); pushIdx += 1 {
		stack = append(stack, pushed[pushIdx])

		stackIdx := len(stack) - 1
		for stackIdx >= 0 && stack[stackIdx] == popped[popIdx] {
			stackIdx -= 1
			popIdx += 1
		}

		stack = stack[:stackIdx+1]
	}

	return len(stack) == 0
}

// @lc code=end

