/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	targetIdx := -1
	leftIdx := 0
	rightIdx := len(nums) - 1

	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if nums[midIdx] == target {
			return midIdx
		}

		if nums[midIdx] >= nums[leftIdx] {
			// rotation at right
			if nums[leftIdx] <= target && nums[midIdx] > target {
				rightIdx = midIdx - 1
			} else {
				leftIdx = midIdx + 1
			}
		} else {
			// rotation at left
			if nums[midIdx] < target && nums[rightIdx] >= target {
				leftIdx = midIdx + 1
			} else {
				rightIdx = midIdx - 1
			}
		}
	}

	return targetIdx
}

// @lc code=end

