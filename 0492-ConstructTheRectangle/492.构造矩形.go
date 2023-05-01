/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func constructRectangle(area int) []int {
	// solution: 数学推导
	//			 W <= L && W * L = area
	//			 W * W <= W * L <= area
	//			 W <= sqrt(area) && area % W == 0
	w := int(math.Sqrt(float64(area)))
	for ; area%w != 0; w -= 1 {

	}

	return []int{area / w, w}
}

// @lc code=end

