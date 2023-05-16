/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	// solution1: 堆排序
	// solution2: 快排
	k = len(nums) - k
	nums = quickSort(nums, 0, len(nums), k)
	return nums[k]
}

func quickSort(nums []int, leftIdx, rightIdx, k int) []int {
	if leftIdx < rightIdx {
		pivotIdx := partition(nums, leftIdx, rightIdx)
		if pivotIdx > k {
			nums = quickSort(nums, leftIdx, pivotIdx, k)
		} else if pivotIdx < k {
			nums = quickSort(nums, pivotIdx+1, rightIdx, k)
		} else {
			return nums
		}
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

