/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */

// @lc code=start
func canWinNim(n int) bool {
	// solution: 每次拿1-3颗，则每轮次至少拿4颗，做为先手保证拿掉余数即可获胜
	return n%4 > 0
}

// @lc code=end

