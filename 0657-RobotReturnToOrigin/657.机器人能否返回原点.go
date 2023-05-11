/*
 * @lc app=leetcode.cn id=657 lang=golang
 *
 * [657] 机器人能否返回原点
 */

// @lc code=start
func judgeCircle(moves string) bool {
	// solution: 模拟
	x, y := 0, 0
	for _, move := range moves {
		switch move {
		case 'U':
			x -= 1
		case 'D':
			x += 1
		case 'L':
			y -= 1
		case 'R':
			y += 1
		}

	}

	return x == 0 && y == 0
}

// @lc code=end

