/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	findSign := true
	isMinus := false
	ans := 0
	for _, char := range s {
		if findSign {
			if char == ' ' {
				continue
			} else if char == '-' {
				isMinus = true
				findSign = false
				continue
			} else if char == '+' {
				findSign = false
				continue
			} else if char >= '0' && char <= '9' {
				findSign = false
				ans = ans*10 + int(char-'0')
				continue
			} else {
				break
			}
		}

		if char < '0' || char > '9' {
			break
		}

		ans = ans*10 + int(char-'0')

		if ans >= math.MaxInt32 && !isMinus {
			ans = math.MaxInt32
			break
		} else if ans > math.MaxInt32 && isMinus {
			ans = math.MinInt32
			break
		}
	}

	if isMinus && ans > 0 {
		ans = -ans
	}

	return ans
}

// @lc code=end

