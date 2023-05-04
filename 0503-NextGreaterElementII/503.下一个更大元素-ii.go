/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	// solution: 单调栈 + 循环数组，因为是循环数组，所以要遍历左边及右边的值，可以进行数组拼接解决这个问题
	res := make([]int, len(nums))
	for idx := range res {
		res[idx] = -1
	}

	numIdxStack := make([]int, 0, len(nums))
	for idx := 0; idx < 2*len(nums)-1; idx += 1 {
		for len(numIdxStack) > 0 && nums[idx%len(nums)] > nums[numIdxStack[len(numIdxStack)-1]] {
			res[numIdxStack[len(numIdxStack)-1]] = nums[idx%len(nums)]
			numIdxStack = numIdxStack[:len(numIdxStack)-1]
		}

		numIdxStack = append(numIdxStack, idx%len(nums))
	}

	return res
}

// @lc code=end

