/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	// solution: 从右边遍历找出nums[idx-1] < nums[idx]，则nums[idx+1]到nums[len(nums)-1]一定是降序
	//			 从降序中找出大于nums[idx-1]的最小数，与nums[idx-1]交换，再将nums[idx+1]到nums[len(nums)-1]的数字头尾互换即可
	//			 example: [1, 2, 3]				idx-1 = 1		[1, 3, 2]
	//					  [3, 2, 1]				idx-1 = -1		[1, 2, 3]
	//					  [1, 2, 3, 9, 8]		idx-1 = 2		[1, 2, 8, 3, 9]
	//					  [1, 2, 3, 9, 8, 7]	idx-1 = 2		[1, 2, 7, 3, 8, 9]
	if len(nums) <= 1 {
		return
	}
	targetIdx := -1
	for idx := len(nums) - 1; idx > 0; idx -= 1 {
		if nums[idx-1] < nums[idx] {
			targetIdx = idx - 1
			break
		}
	}

	if targetIdx >= 0 {
		for idx := len(nums) - 1; idx > targetIdx; idx -= 1 {
			if nums[idx] > nums[targetIdx] {
				nums[targetIdx], nums[idx] = nums[idx], nums[targetIdx]
				break
			}
		}
	}

	for leftIdx, rightIdx := targetIdx+1, len(nums)-1; leftIdx < rightIdx; {
		nums[leftIdx], nums[rightIdx] = nums[rightIdx], nums[leftIdx]
		leftIdx += 1
		rightIdx -= 1
	}
}

// @lc code=end

