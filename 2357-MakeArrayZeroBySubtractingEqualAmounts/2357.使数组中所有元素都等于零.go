/*
 * @lc app=leetcode.cn id=2357 lang=golang
 *
 * [2357] 使数组中所有元素都等于零
 */

// @lc code=start
func minimumOperations(nums []int) int {
	// solution1: 排序，然后遍历数组
	/**
	sort.Ints(nums)

	minusNum, counter := 0, 0

	for _, num := range nums {
		if num == 0 {
			continue
		}

		if num-minusNum != 0 {
			minusNum += (num - minusNum)
			counter += 1
		}
	}

	return counter
	*/

	// solution2: hash表存放非零值
	numMap := make(map[int]struct{}, 0)
	for _, num := range nums {
		if num == 0 {
			continue
		}

		numMap[num] = struct{}{}
	}

	return len(numMap)
}

// @lc code=end

