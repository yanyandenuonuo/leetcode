/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

// @lc code=start
func calculate(s string) int {
	// solution: 利用栈，因为只有加减，所以移除所有括号后影响的只有数字前的符号
	//			 利用栈来记录符号变化即可1表示+，-1表示-
	//			 遇到(则将最后一个sign入栈
	//			 遇到)出栈
	//			 遇到+ - 则使用栈顶元素更新sign值
	signStack := []int{1}
	sign := 1 // 默认为+
	res := 0
	for idx, isValid := 0, false; idx < len(s); {
		switch s[idx] {
		case ' ':
			idx += 1
		case '-':
			isValid = true
			sign = -signStack[len(signStack)-1]
			idx += 1
		case '+':
			if !isValid {
				return 0
			}
			sign = signStack[len(signStack)-1]
			idx += 1
		case '(':
			isValid = true
			signStack = append(signStack, sign)
			idx += 1
		case ')':
			isValid = true
			signStack = signStack[:len(signStack)-1]
			idx += 1
		default:
			isValid = true
			num := 0
			for ; idx < len(s) && s[idx] >= '0' && s[idx] <= '9'; idx += 1 {
				num = num*10 + int(s[idx]-'0')
			}

			res += sign * num
		}
	}
	return res
}

// @lc code=end

