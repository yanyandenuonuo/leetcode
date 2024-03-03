/*
 * @lc app=leetcode.cn id=2455 lang=golang
 *
 * [2455] 可被三整除的偶数的平均值
 */

// @lc code=start
func averageValue(nums []int) int {
	total, counter := 0, 0
	for _, val := range nums {
		if val&0x01 == 0x00 && val%3 == 0 {
			total += val
			counter += 1
		}
	}

	if counter == 0 {
		return 0
	}

	return total / counter
}

// @lc code=end

