/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	res := "1"
	for idx := 2; idx <= n; idx += 1 {
		nextRes := ""
		for resIdx := len(res) - 1; resIdx >= 0; {
			// find repeat
			repeatCount := 1
			for repeatIdx := resIdx - 1; repeatIdx >= 0; repeatIdx -= 1 {
				if res[repeatIdx] != res[resIdx] {
					break
				}
				repeatCount += 1
			}

			nextRes = strconv.Itoa(repeatCount) + string(res[resIdx]) + nextRes
			resIdx -= repeatCount
		}
		res = nextRes
	}
	return res
}

// @lc code=end

