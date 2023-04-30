/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	// solution1: 利用位排序
	// solution2: hash
	// solution3: 位运算，在数组后添加[0, n]，则新数组中只有1个数出现了依次，利用异或取结果即可
	// solution4: 求和, [0, n]的sum = n * (n+1) / 2，差值即为缺失的数字
	sum := 0
	for _, val := range nums {
		sum += val
	}

	return len(nums)*(len(nums)+1)/2 - sum
}

// @lc code=end

