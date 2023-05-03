/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	// solution1: 直接对平方数排序
	// solution2: 找出正负数分界点，然后双指针移动，正序排序平方数
	/**
	negativeIdx, idx := -1, 0
	for idx < len(nums) {
		if nums[idx] < 0 {
			negativeIdx = idx
		}
		if idx != negativeIdx {
			break
		}
		idx += 1
	}
	squareList := make([]int, 0, len(nums))
	for negativeIdx >= 0 || idx < len(nums) {
		if negativeIdx >= 0 && idx < len(nums) {
			if nums[negativeIdx]*nums[negativeIdx] < nums[idx]*nums[idx] {
				squareList = append(squareList, nums[negativeIdx]*nums[negativeIdx])
				negativeIdx -= 1
			} else {
				squareList = append(squareList, nums[idx]*nums[idx])
				idx += 1
			}

		} else if negativeIdx >= 0 {
			squareList = append(squareList, nums[negativeIdx]*nums[negativeIdx])
			negativeIdx -= 1
		} else if idx < len(nums) {
			squareList = append(squareList, nums[idx]*nums[idx])
			idx += 1
		}
	}

	return squareList
	*/

	// solution3: 头尾双指针，逆序排序平方数
	res := make([]int, len(nums))
	for leftIdx, rightIdx, resIdx := 0, len(nums)-1, len(nums)-1; leftIdx <= rightIdx; resIdx -= 1 {
		leftSquare := nums[leftIdx] * nums[leftIdx]
		rightSquare := nums[rightIdx] * nums[rightIdx]
		if leftSquare > rightSquare {
			res[resIdx] = leftSquare
			leftIdx += 1
		} else {
			res[resIdx] = rightSquare
			rightIdx -= 1
		}
	}

	return res
}

// @lc code=end

