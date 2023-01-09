/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
func rand10() int {
	randNum := (rand7()-1)*7 + rand7()
	for randNum > 40 {
		randNum = (rand7()-1)*7 + rand7()
	}
	return randNum%10 + 1
}

// @lc code=end

