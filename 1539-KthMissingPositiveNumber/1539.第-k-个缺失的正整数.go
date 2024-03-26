/*
 * @lc app=leetcode.cn id=1539 lang=golang
 *
 * [1539] 第 k 个缺失的正整数
 */

// @lc code=start
func findKthPositive(arr []int, k int) int {
	// solution1: 位标识
	// solution2: 缺失值统计
	// 			  k在数组之前
	// 			  k在数组之后
	// 			  k在数组之间
	if arr[0]-1 >= k {
		return k
	} else if (arr[len(arr)-1] - len(arr)) < k {
		return k + len(arr)
	} else {
		missCounter, preVal := arr[0]-1, arr[0]-1
		for _, val := range arr {
			missCounter += (val - preVal - 1)
			if missCounter < k {
				preVal = val
				continue
			}

			preVal = val - 1
			for missCounter > k {

				preVal -= 1
				missCounter -= 1
			}

			return preVal
		}

		return -1
	}

	// solution3: 二分查找所在区间，对solution2再次优化
}

// @lc code=end

