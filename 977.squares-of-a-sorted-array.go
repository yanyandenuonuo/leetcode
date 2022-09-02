/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	negativeIdx := -1
	idx := 0
	numsLength := len(nums)
	for idx < numsLength {
		if nums[idx] < 0 {
			negativeIdx = idx
		}
		if idx != negativeIdx {
			break;
		}
		idx += 1
	}
	squareList := make([]int, 0, numsLength)
	for negativeIdx >= 0 || idx < numsLength {
		if negativeIdx >= 0 && idx < numsLength {
			if nums[negativeIdx] * nums[negativeIdx] < nums[idx] * nums[idx] {
				squareList = append(squareList, nums[negativeIdx] * nums[negativeIdx])
				negativeIdx -= 1
			} else {
				squareList = append(squareList, nums[idx] * nums[idx])
				idx += 1
			}

		} else if negativeIdx >= 0 {
			squareList = append(squareList, nums[negativeIdx] * nums[negativeIdx])
				negativeIdx -= 1
		} else if idx < numsLength {
			squareList = append(squareList, nums[idx] * nums[idx])
				idx += 1
		}
	}
	return squareList
}
// @lc code=end

