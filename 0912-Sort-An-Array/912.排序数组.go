/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	return quickSort(nums, 0, len(nums))
}

func quickSort(nums []int, leftIdx, rightIdx int) []int {
	if leftIdx < rightIdx {
		pivot := partition(nums, leftIdx, rightIdx)
		nums = quickSort(nums, leftIdx, pivot)
		nums = quickSort(nums, pivot+1, rightIdx)
	}
	return nums
}

func partition(nums []int, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx)
	// swap pivot with right-1
	nums[rightIdx-1], nums[pivotIdx] = nums[pivotIdx], nums[rightIdx-1]
	pivotIdx = rightIdx - 1

	swapIdx := leftIdx
	for idx := leftIdx; idx < rightIdx-1; idx += 1 {
		if nums[idx] < nums[pivotIdx] {
			nums[idx], nums[swapIdx] = nums[swapIdx], nums[idx]
			swapIdx += 1
		}
	}
	nums[swapIdx], nums[pivotIdx] = nums[pivotIdx], nums[swapIdx]

	return swapIdx
}

// @lc code=end

