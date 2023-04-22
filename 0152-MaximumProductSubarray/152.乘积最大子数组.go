/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	// solution1: 滑动窗口
	// solution2: 动态规划
	//			  dp[i][0]表示以i结尾的子数组最大乘积
	//			  dp[i][1]表示以i结尾的子数组最小乘积，因为存在偶数个负数相乘得到正数的场景
	// 因为是乘积，所以下面三个初始值均设置位nums[0]
	preMaxVal, preMinVal, maxVal := nums[0], nums[0], nums[0]

	for idx := 1; idx < len(nums); idx += 1 {
		// 作为子数组开头
		currentMax := nums[idx]

		// 作为子数组结尾，正数场景
		maxVal1 := preMaxVal * nums[idx]
		if maxVal1 > currentMax {
			currentMax = maxVal1
		}

		// 作为子数组结尾，负数场景
		maxVal2 := preMinVal * nums[idx]
		if maxVal2 > currentMax {
			currentMax = maxVal2
		}

		// 作为子数组开头
		currentMin := nums[idx]

		// 作为子数组结尾，负数场景
		minVal1 := preMinVal * nums[idx]
		if minVal1 < currentMin {
			currentMin = minVal1
		}

		// 作为子数组结尾，正数场景
		minVal2 := preMaxVal * nums[idx]
		if minVal2 < currentMin {
			currentMin = minVal2
		}

		preMaxVal = currentMax
		preMinVal = currentMin

		if currentMax > maxVal {
			maxVal = currentMax
		}
	}
	return maxVal
}

// @lc code=end

