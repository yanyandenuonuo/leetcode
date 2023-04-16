/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort
	quickSort(intervals, 0, len(intervals)-1)

	// merge
	res := make([][]int, 0, len(intervals))

	interval := intervals[0]

	for idx := 1; idx < len(intervals); idx += 1 {
		if interval[1] < intervals[idx][0] {
			// [     ]
			//          [     ]
			res = append(res, interval)
			interval = intervals[idx]
		} else if intervals[idx][0] <= interval[1] && interval[1] <= intervals[idx][1] {
			// [     ]
			//    [     ]
			interval[1] = intervals[idx][1]
		} else {
			// never exist because the intervals had been sorted
			// [     ]
			//  [ ]
			// ...
		}
	}

	res = append(res, interval)

	return res
}

func quickSort(nums [][]int, leftIdx, rightIdx int) {
	if leftIdx >= rightIdx {
		return
	}

	pivot := findPivotIdx(nums, leftIdx, rightIdx)
	quickSort(nums, leftIdx, pivot-1)
	quickSort(nums, pivot+1, rightIdx)
}

func findPivotIdx(nums [][]int, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx+1)

	nums[pivotIdx], nums[rightIdx] = nums[rightIdx], nums[pivotIdx]
	swapIdx := leftIdx
	for idx := leftIdx; idx < rightIdx; idx += 1 {
		if nums[idx][0] < nums[rightIdx][0] {
			nums[swapIdx], nums[idx] = nums[idx], nums[swapIdx]
			swapIdx += 1
		}
	}

	nums[swapIdx], nums[rightIdx] = nums[rightIdx], nums[swapIdx]
	return swapIdx
}

// @lc code=end

