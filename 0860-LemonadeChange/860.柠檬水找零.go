/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			five += 1
		case 10:
			if five < 1 {
				return false
			}

			ten += 1
			five -= 1
		case 20:
			if ten < 1 {
				if five < 3 {
					return false
				}

				five -= 3
			} else {
				if five < 1 {
					return false
				}

				ten -= 1
				five -= 1
			}
		default:
			return false
		}
	}

	return true
}

// @lc code=end

