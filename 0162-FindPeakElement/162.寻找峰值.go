/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	// solution1: 遍历
	/**
	if len(nums) == 1 {
		return 0
	}

	leftIdx, rightIdx := 0, len(nums)-1
	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2

		preVal := 0
		if midIdx-1 < 0 {
			preVal = nums[midIdx] - 1
		} else {
			preVal = nums[midIdx-1]
		}

		nextVal := 0
		if midIdx+1 < len(nums) {
			nextVal = nums[midIdx+1]
		} else {
			nextVal = nums[midIdx] - 1
		}

		if preVal < nums[midIdx] && nums[midIdx] > nextVal {
			return midIdx
		}
		if nums[midIdx] < nums[midIdx+1] {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}
	}
	return -1
	*/

	// solution2: 二分查找
	return searchPeakElement(nums, 0, len(nums)-1)
}

func searchPeakElement(nums []int, leftIdx, rightIdx int) int {
	if leftIdx > rightIdx {
		return -1
	}
	randIdx := leftIdx + rand.Intn(rightIdx-leftIdx+1)
	if (randIdx-1 < 0 || nums[randIdx-1] < nums[randIdx]) && (randIdx+1 > len(nums)-1 || nums[randIdx] > nums[randIdx+1]) {
		return randIdx
	}

	targetIdx := searchPeakElement(nums, leftIdx, randIdx-1)
	if targetIdx != -1 {
		return targetIdx
	}

	targetIdx = searchPeakElement(nums, randIdx+1, rightIdx)
	if targetIdx != -1 {
		return targetIdx
	}
	return -1
}

// @lc code=end

