/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	res[0] = findLeftTarget(nums, target)
	res[1] = findRightTarget(nums, target)

	return res
}

func findLeftTarget(nums []int, target int) int {
	targetIdx := -1
	leftIdx := 0
	rightIdx := len(nums) - 1
	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if nums[midIdx] > target {
			rightIdx = midIdx - 1
		} else if nums[midIdx] == target {
			targetIdx = midIdx
			rightIdx = midIdx - 1
		} else {
			leftIdx = midIdx + 1
		}
	}

	return targetIdx
}

func findRightTarget(nums []int, target int) int {
	targetIdx := -1
	leftIdx := 0
	rightIdx := len(nums) - 1
	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if nums[midIdx] > target {
			rightIdx = midIdx - 1
		} else if nums[midIdx] == target {
			targetIdx = midIdx
			leftIdx = midIdx + 1
		} else {
			leftIdx = midIdx + 1
		}
	}

	return targetIdx
}

// @lc code=end

