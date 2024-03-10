/*
 * @lc app=leetcode.cn id=2447 lang=golang
 *
 * [2447] 最大公因数等于 K 的子数组数目
 */

// @lc code=start
func subarrayGCD(nums []int, k int) int {
	// solution1: 找出子数组元素与k的最大公约数为k的index，再求子数组个数
	// solution2: 动态规划dp[i][j]表示子数组nums[i-j]的最大公约数，最后求k的个数即可
	//			  可以剪枝若该数组元素与k的最大公约数不为k，则可跳过
	counter := 0
	for i := 0; i < len(nums); i += 1 {
		gcd := nums[i]

		for j := i; j < len(nums); j += 1 {
			gcd = findGCD([]int{gcd, nums[j]})
			if gcd == k {
				counter += 1
			} else if gcd < k {
				break
			}
		}
	}

	return counter
}

func findGCD(nums []int) int {
	minVal, maxVal := nums[0], nums[0]

	for _, val := range nums {
		if val > maxVal {
			maxVal = val
		} else if val < minVal {
			minVal = val
		}
	}

	for minVal != 0 {
		minVal, maxVal = maxVal%minVal, minVal
	}

	return maxVal
}

// @lc code=end

