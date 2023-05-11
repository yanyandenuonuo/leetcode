/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] 棒球比赛
 */

// @lc code=start
func calPoints(operations []string) int {
	// solution: 栈
	score, scoreStack := 0, make([]int, 0, len(operations))
	for _, ops := range operations {
		switch ops {
		case "+":
			x := scoreStack[len(scoreStack)-2] + scoreStack[len(scoreStack)-1]
			score += x
			scoreStack = append(scoreStack, x)
		case "D":
			x := scoreStack[len(scoreStack)-1] * 2
			score += x
			scoreStack = append(scoreStack, x)
		case "C":
			score -= scoreStack[len(scoreStack)-1]
			scoreStack = scoreStack[:len(scoreStack)-1]
		default:
			x, _ := strconv.Atoi(ops)
			score += x
			scoreStack = append(scoreStack, x)
		}
	}

	return score
}

// @lc code=end

