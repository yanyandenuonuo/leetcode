/*
 * @lc app=leetcode.cn id=1748 lang=golang
 *
 * [1748] 唯一元素的和
 */

// @lc code=start
func sumOfUnique(nums []int) int {
	numMap, sum := make(map[int]int), 0
	for _, num := range nums {
		numMap[num] += 1
		switch numMap[num] {
		case 1:
			sum += num
		case 2:
			sum -= num
		}
	}

	return sum
}

// @lc code=end

