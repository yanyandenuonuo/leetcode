/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	// solution1: 递归
	// solution2: 栈
	numStack, stringStack := make([]int, 0), make([]string, 0)
	for idx := 0; idx < len(s); {
		if s[idx] >= '0' && s[idx] <= '9' {
			num := 0
			for idx < len(s) && s[idx] >= '0' && s[idx] <= '9' {
				num = num*10 + int(s[idx]-'0')
				idx += 1
			}
			numStack = append(numStack, num)
		} else if s[idx] >= 'a' && s[idx] <= 'z' {
			sBuilder := new(strings.Builder)
			for idx < len(s) && s[idx] >= 'a' && s[idx] <= 'z' {
				sBuilder.WriteByte(s[idx])
				idx += 1
			}
			stringStack = append(stringStack, sBuilder.String())
		} else if s[idx] == '[' {
			stringStack = append(stringStack, "[")
			idx += 1
		} else if s[idx] == ']' {
			// 找出同一级别要拼接的
			strIdx, currentStr := len(stringStack)-1, ""
			for ; strIdx >= 0 && stringStack[strIdx] != "["; strIdx -= 1 {
				currentStr = stringStack[strIdx] + currentStr
			}
			stringStack = stringStack[:strIdx]

			// 与数字结合
			sBuilder := new(strings.Builder)
			for counter := 0; counter < numStack[len(numStack)-1]; counter += 1 {
				sBuilder.WriteString(currentStr)
			}
			stringStack = append(stringStack, sBuilder.String())

			numStack = numStack[:len(numStack)-1]
			idx += 1
		}
	}

	sBuilder := new(strings.Builder)
	for _, str := range stringStack {
		sBuilder.WriteString(str)
	}

	return sBuilder.String()
}

// @lc code=end

