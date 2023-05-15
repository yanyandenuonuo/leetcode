/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	// solution1: 模拟
	numLength := len(nums)
	if k > numLength {
		k = k % numLength
	}

	previousLength := numLength - k
	if previousLength <= 0 {
		return
	}

	if k > (numLength - k) {
		// previous is less, use extra storage save previous data
		previousNums := make([]int, previousLength)
		for idx := 0; idx < previousLength; idx += 1 {
			previousNums[idx] = nums[idx]
		}

		for idx := 0; idx < k; idx += 1 {
			nums[idx] = nums[numLength-k+idx]
		}

		for idx := k; idx < numLength; idx += 1 {
			nums[idx] = previousNums[idx-k]
		}
	} else {
		// last is more, use extra storage save last data
		lastNums := make([]int, k)
		for idx := 0; idx < k; idx += 1 {
			lastNums[idx] = nums[numLength-k+idx]
		}

		for idx := 0; idx < previousLength; idx += 1 {
			nums[numLength-1-idx] = nums[numLength-1-idx-k]
		}
		for idx := 0; idx < k; idx += 1 {
			nums[idx] = lastNums[idx]
		}
	}

	// solution2: 环状替换
	// todo
	// solution3: 数组翻转，然后前k个反再翻转，后len(nums)-k个再翻转
	// todo
}

// @lc code=end

