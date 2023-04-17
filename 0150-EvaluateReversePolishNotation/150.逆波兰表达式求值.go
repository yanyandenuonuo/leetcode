/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		switch token {
		case "+":
			stack = append(stack[0:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
		case "-":
			stack = append(stack[0:len(stack)-2], stack[len(stack)-2]-stack[len(stack)-1])
		case "*":
			stack = append(stack[0:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
		case "/":
			stack = append(stack[0:len(stack)-2], stack[len(stack)-2]/stack[len(stack)-1])
		default:
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}
	return stack[0]
}

// @lc code=end

