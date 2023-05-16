/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 */

// @lc code=start
func asteroidCollision(asteroids []int) []int {
	// solution: 栈，注意只需要考虑正向碰上逆向的问题，也不需要考虑环形碰撞问题
	asteroidStack := make([]int, 0)
	for _, asteroid := range asteroids {
		destroy := false
		for len(asteroidStack) > 0 && asteroidStack[len(asteroidStack)-1] > 0 && asteroid < 0 {
			if asteroidStack[len(asteroidStack)-1] < -asteroid {
				asteroidStack = asteroidStack[:len(asteroidStack)-1]
			} else if asteroidStack[len(asteroidStack)-1] > -asteroid {
				destroy = true
				break
			} else {
				asteroidStack = asteroidStack[:len(asteroidStack)-1]
				destroy = true
				break
			}
		}

		if !destroy {
			asteroidStack = append(asteroidStack, asteroid)
		}
	}

	return asteroidStack
}

// @lc code=end

