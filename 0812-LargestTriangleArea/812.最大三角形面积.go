/*
 * @lc app=leetcode.cn id=812 lang=golang
 *
 * [812] 最大三角形面积
 */

// @lc code=start
func largestTriangleArea(points [][]int) float64 {
	// solution1: 枚举，三角形三个顶点坐标为(x1, y1), (x2, y2), (x3, y3)，则面积s = |x1*y2 + x2*y3 + x3*y1- x1*y3- x2*y1 - x3*y2|/2
	res := float64(0)
	for idx1 := 2; idx1 < len(points); idx1 += 1 {
		for idx2 := 1; idx2 < idx1; idx2 += 1 {
			for idx3 := 0; idx3 < idx2; idx3 += 1 {
				triangleArea := getTriangleArea(points[idx1][0], points[idx1][1], points[idx2][0], points[idx2][1], points[idx3][0], points[idx3][1])
				if triangleArea > res {
					res = triangleArea
				}
			}
		}
	}

	return res

	// solution2: 凸包
	// todo
}

func getTriangleArea(x1, y1, x2, y2, x3, y3 int) float64 {
	return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
}

// @lc code=end

