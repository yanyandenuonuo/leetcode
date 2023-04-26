/*
 * @lc app=leetcode.cn id=640 lang=golang
 *
 * [640] 求解方程
 */

// @lc code=start
func solveEquation(equation string) string {
	// solution: 合并项，最终行如numX * X + numConst = 0
	numX, numConst, signStack, sign, signIdx := 0, 0, []int{1, -1}, 1, 0
	for idx := 0; idx < len(equation); {
		switch equation[idx] {
		case ' ':
			idx += 1
		case '=':
			signIdx += 1
			sign = signStack[signIdx]
			idx += 1
		case '+':
			sign = signStack[signIdx]
			idx += 1
		case '-':
			sign = -signStack[signIdx]
			idx += 1
		case 'x':
			numX += sign
			idx += 1
		default:
			num := 0
			for ; idx < len(equation) && equation[idx] >= '0' && equation[idx] <= '9'; idx += 1 {
				num = num*10 + int(equation[idx]-'0')
			}
			if idx < len(equation) && equation[idx] == 'x' {
				numX += sign * num
				idx += 1
				continue
			}
			numConst += sign * num
		}
	}
	if numX == 0 {
		if numConst == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}
	return "x=" + strconv.Itoa(-numConst/numX)
}

// @lc code=end

