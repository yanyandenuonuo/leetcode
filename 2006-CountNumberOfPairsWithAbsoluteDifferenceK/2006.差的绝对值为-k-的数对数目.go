/*
 * @lc app=leetcode.cn id=2006 lang=golang
 *
 * [2006] 差的绝对值为 K 的数对数目
 */

// @lc code=start
func countKDifference(nums []int, k int) int {
	// solution1: 枚举
	// solution2: hash表计数，然后再遍历nums并更新hash表中计数值
	numMap := make(map[int]int, len(nums))
	for _, val := range nums {
		numMap[val] += 1
	}

	res := 0
	for _, val := range nums {
		res += numMap[val-k] + numMap[val+k]
		numMap[val] -= 1
	}

	return res
}

// @lc code=end

