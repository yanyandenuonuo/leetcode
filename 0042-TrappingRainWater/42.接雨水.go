/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// solution1: 单调栈
	/**
	stack := make([]int, 0, len(height))
	sum := 0
	for idx := range height {
		for len(stack) > 0 && height[idx] > height[stack[len(stack)-1]] {
			if len(stack) == 1 {
				stack = stack[:len(stack)-1]
				break
			}

			topIdx := stack[len(stack)-1]

			stack = stack[:len(stack)-1]

			leftIdx := stack[len(stack)-1]

			minHeight := height[idx]
			if height[leftIdx] < minHeight {
				minHeight = height[leftIdx]
			}
			sum += (idx - leftIdx - 1) * (minHeight - height[topIdx])
		}
		stack = append(stack, idx)
	}
	return sum
	*/

	// solution2: 动态规划，依次求出idx处左边的最高边界以及右边的最高边界，雨水量即为min(leftMax, rightMax)-height[idx]
	/**
	heightNums := make([]int, len(height))
	// leftMax
	heightNums[0] = height[0]
	for idx := 1; idx < len(heightNums); idx += 1 {
		if heightNums[idx-1] > height[idx] {
			heightNums[idx] = heightNums[idx-1]
		} else {
			heightNums[idx] = height[idx]
		}
	}
	// rightMax
	if heightNums[len(heightNums)-1] > height[len(height)-1] {
		heightNums[len(heightNums)-1] = height[len(height)-1]
	}

	for idx := len(heightNums) - 2; idx >= 0; idx -= 1 {
		rightMax := height[idx]
		if heightNums[idx+1] > rightMax {
			rightMax = heightNums[idx+1]
		}
		if rightMax < heightNums[idx] {
			heightNums[idx] = rightMax
		}
	}
	sum := 0
	for idx, heightMax := range heightNums {
		sum += heightMax - height[idx]
	}
	return sum
	*/

	// solution3: 双指针，实质为solution2的优化解决方案
	sum := 0
	// 因为每次调整的idx都是较小的idx，所以最后leftIdx == rightIdx时一定是最高的点，最高点无法存水，所以可以不用特殊考虑
	for leftIdx, rightIdx, leftMax, rightMax := 0, len(height)-1, 0, 0; leftIdx < rightIdx; {
		if height[leftIdx] > leftMax {
			leftMax = height[leftIdx]
		}
		if height[rightIdx] > rightMax {
			rightMax = height[rightIdx]
		}
		if height[leftIdx] < height[rightIdx] {
			sum += leftMax - height[leftIdx]
			leftIdx += 1
		} else {
			sum += rightMax - height[rightIdx]
			rightIdx -= 1
		}
	}
	return sum
}

// @lc code=end

