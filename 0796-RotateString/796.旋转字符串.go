/*
 * @lc app=leetcode.cn id=796 lang=golang
 *
 * [796] 旋转字符串
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	// s1 = xy, s2 = yx, s1s1 = xyxy = xs2y
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}

// @lc code=end

