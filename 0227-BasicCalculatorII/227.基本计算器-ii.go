/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
func calculate(s string) int {
	// solution: 利用栈，先消除乘除运算，然后对栈求和即可
	preSign := "+"
	numStack := make([]int, 0, len(s))
	for idx := 0; idx < len(s); {
		switch s[idx] {
		case ' ':
			idx += 1
		case '+', '-', '*', '/':
			preSign = string(s[idx])
			idx += 1
		default:
			num := 0
			for ; idx < len(s) && s[idx] >= '0' && s[idx] <= '9'; idx += 1 {
				num = num*10 + int(s[idx]-'0')
			}
			switch preSign {
			case "+":
				numStack = append(numStack, num)
			case "-":
				numStack = append(numStack, -num)
			case "*":
				numStack[len(numStack)-1] = numStack[len(numStack)-1] * num
			case "/":
				numStack[len(numStack)-1] = numStack[len(numStack)-1] / num
			}
		}
	}
	res := 0
	for idx := len(numStack) - 1; idx >= 0; idx -= 1 {
		res += numStack[idx]
	}
	return res
}

// @lc code=end

