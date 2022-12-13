/*
 * @lc app=leetcode.cn id=2195 lang=golang
 *
 * [2195] 向数组中追加 K 个整数
 */

// @lc code=start
func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)

	res, previousValue, counter := int64(0), 0, 0

	for numIdx := 0; counter < k && numIdx < len(nums); numIdx += 1 {
		if nums[numIdx]-previousValue > 1 && counter < k {
			numCount := nums[numIdx] - previousValue - 1
			if numCount > k-counter {
				break
			}

			sum := previousValue + nums[numIdx]

			res += int64((sum) / 2 * numCount)
			if sum%2 == 1 {
				res += int64(numCount / 2)
			}

			counter += numCount
		}
		previousValue = nums[numIdx]
	}

	if counter < k {
		numCount := k - counter
		sum := (previousValue + 1) + (previousValue + numCount)

		res += int64((sum) / 2 * numCount)
		if sum%2 == 1 {
			res += int64(numCount / 2)
		}
	}

	return res
}

// @lc code=end

