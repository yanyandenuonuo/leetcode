/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
	// solution1: 排序
	//			  全负数|全正数，最大的三个整数乘积
	//			  部分正+部分负，三个最大正数乘积或两个最小负数和一个最大正数乘积
	// solution2: 直接遍历找5个数，3个最大，2个最小
	max1, max2, max3, min1, min2 := -1<<31, -1<<31, -1<<31, 1<<31-1, 1<<31-1

	for _, num := range nums {
		if num > max1 {
			max1, max2, max3 = num, max1, max2
		} else if num > max2 {
			max2, max3 = num, max2
		} else if num > max3 {
			max3 = num
		}

		if num < min1 {
			min1, min2 = num, min1
		} else if num < min2 {
			min2 = num
		}
	}

	if max1*max2*max3 > min1*min2*max1 {
		return max1 * max2 * max3
	}

	return min1 * min2 * max1
}

// @lc code=end

