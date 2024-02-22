/*
 * @lc app=leetcode.cn id=836 lang=golang
 *
 * [836] 矩形重叠
 */

// @lc code=start
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// solution1: 判断重合区域

	/**
	if rec2[0] <= rec1[0] && rec1[0] < rec2[2] &&
		rec2[1] <= rec1[1] && rec1[1] < rec2[3] &&
		rec2[0] < rec1[2] && rec1[2] <= rec2[2] &&
		rec2[1] < rec1[3] && rec1[3] <= rec2[3] {
		// rec1小于或等于rec2
		return true
	} else if rec1[0] <= rec2[0] && rec2[0] < rec1[2] &&
		rec1[1] <= rec2[1] && rec2[1] < rec1[3] {
		// rec2左下角在rec1内
		return true
	} else if rec1[0] <= rec2[0] && rec2[0] < rec1[2] &&
		rec1[1] < rec2[3] && rec2[3] <= rec1[3] {
		// rec2左上角在rec1内
		return true
	} else if rec1[0] < rec2[2] && rec2[2] <= rec1[2] &&
		rec1[1] < rec2[3] && rec2[3] <= rec1[3] {
		// rec2右上角在rec1内
		return true
	} else if rec1[0] < rec2[2] && rec2[2] <= rec1[2] &&
		rec1[1] <= rec2[1] && rec2[1] < rec1[3] {
		// rec2右下角在rec1内
		return true
	} else if rec2[1] <= rec1[1] && rec2[3] >= rec1[3] &&
		((rec2[0] >= rec1[0] && rec2[0] < rec1[2]) ||
			(rec2[2] > rec1[0] && rec2[2] <= rec1[2])) {
		// 十字形、T字形、倒T子形
		return true
	} else if rec2[0] <= rec1[0] && rec2[2] >= rec1[2] &&
		((rec2[1] >= rec1[1] && rec2[1] < rec1[3]) ||
			(rec2[3] > rec1[1] && rec2[3] <= rec1[3])) {
		// 十字形、左横T字形、右横T子形
		return true
	} else {
		return false
	}
	*/

	// solution2: 判断不重合区域

	if rec2[2] <= rec1[0] {
		// rec2在rec1左边
		return false
	} else if rec2[0] >= rec1[2] {
		//rec2在rec1右边
		return false
	} else if rec2[1] >= rec1[3] {
		//rec2在rec1上边
		return false
	} else if rec2[3] <= rec1[1] {
		//rec2在rec1下边
		return false
	} else {
		return true
	}
}

// @lc code=end

