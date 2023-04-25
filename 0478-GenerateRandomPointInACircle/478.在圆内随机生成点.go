/*
 * @lc app=leetcode.cn id=478 lang=golang
 *
 * [478] 在圆内随机生成点
 */

// @lc code=start
type Solution struct {
	radius   float64
	x_center float64
	y_center float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		radius:   radius,
		x_center: x_center,
		y_center: y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	// solution1: 根据半径生成坐标轴，若点在圆内则直接返回，否则重新生成
	/**
	for {
		randX := rand.Float64()*2 - 1
		randY := rand.Float64()*2 - 1
		if (randX*randX + randY*randY) < 1 {
			return []float64{randX*this.radius + this.x_center, randY*this.radius + this.y_center}
		}
	}
	*/

	// solution2: 生成随机角度，然后生成随机半径的长度
	randR := math.Sqrt(rand.Float64()) // 圆的面积为Pi*R*R，如果直接使用[0,1)随机生成R，则越靠近圆心概率越大
	sin, cos := math.Sincos(rand.Float64() * 2 * math.Pi)
	return []float64{randR*this.radius*cos + this.x_center, randR*this.radius*sin + this.y_center}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
// @lc code=end

