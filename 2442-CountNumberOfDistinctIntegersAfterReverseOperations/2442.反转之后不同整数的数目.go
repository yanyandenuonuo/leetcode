/*
 * @lc app=leetcode.cn id=2442 lang=golang
 *
 * [2442] 反转之后不同整数的数目
 */

// @lc code=start
func countDistinctIntegers(nums []int) int {
	numLength := len(nums)
	numMap := make(map[int]bool, numLength*2)
	for idx := 0; idx < numLength; idx += 1 {
		if _, isExist := numMap[nums[idx]]; !isExist {
			numMap[nums[idx]] = true
		}
		reverseNum := reverse(nums[idx])

		if _, isExist := numMap[reverseNum]; !isExist {
			numMap[reverseNum] = true
		}
	}

	return len(numMap)
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}

	return res
}

// @lc code=end

