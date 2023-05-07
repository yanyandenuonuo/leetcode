/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	// solution: 贪心，维护一个最远可以到达的位置，并更新这个位置的值
	maxIdx := 0
	for idx, num := range nums {
		if idx > maxIdx {
			return false
		}

		if idx+num > maxIdx {
			maxIdx = idx + num

			if maxIdx >= len(nums)-1 {
				return true
			}
		}
	}

	return maxIdx >= len(nums)-1
}

// @lc code=end

