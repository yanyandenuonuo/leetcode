/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// solution: 先用二分查找row，再用二分查找column
	for rowKey, columnItem := range matrix {
		if matrix[rowKey][0] > target {
			return false
		}
		if (rowKey+1) < len(matrix) && matrix[rowKey+1][0] <= target {
			continue
		}

		return searchTarget(columnItem, target)
	}

	return false
}

func searchTarget(nums []int, target int) bool {
	leftIdx := 0
	rightIdx := len(nums) - 1
	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if nums[midIdx] == target {
			return true
		} else if nums[midIdx] < target {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}
	}

	return false
}

// @lc code=end

