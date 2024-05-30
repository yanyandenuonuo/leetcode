/*
 * @lc app=leetcode.cn id=2384 lang=golang
 *
 * [2384] 最大回文数字
 */

// @lc code=start
func largestPalindromic(num string) string {
	numCounter := make([]int, 10)
	for _, numChar := range num {
		numCounter[numChar-'0'] += 1
	}

	midNum := -1

	numList := make([]string, 0, len(num))
	for idx := 9; idx >= 0; {
		if numCounter[idx] <= 0 {
			idx -= 1
		} else if numCounter[idx] == 1 {
			if idx > midNum {
				midNum = idx
			}

			idx -= 1
		} else {
			if idx == 0 && len(numList) == 0 {
				break
			}

			numList = append(numList, string(rune('0'+idx)))
			numCounter[idx] -= 2
		}
	}

	if midNum >= 0 {
		numList = append(numList, string(rune('0'+midNum)))
	} else {
		numList = append(numList, "")
	}

	for idx := len(numList) - 2; idx >= 0; idx -= 1 {
		numList = append(numList, numList[idx])
	}

	res := strings.Join(numList, "")

	if len(res) == 0 && len(num) > 0 {
		return "0"
	}

	return res
}

// @lc code=end

