/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	// solution1: 暴力
	// solution2: 位排序
	// solution3: 排序并记录位置信息
	arr := make([]*numWithIdx, len(nums))

	for idx, num := range nums {
		arr[idx] = &numWithIdx{
			Num: num,
			Idx: idx,
		}
	}

	quickSort(arr, 0, len(arr)-1)

	resList := make([]int, len(nums))

	for idx, val := range arr {
		if idx > 0 && val.Num == arr[idx-1].Num {
			resList[val.Idx] = resList[arr[idx-1].Idx]
		} else {
			resList[val.Idx] = idx
		}
	}

	return resList
}

type numWithIdx struct {
	Num int
	Idx int
}

func quickSort(arr []*numWithIdx, leftIdx, rightIdx int) {
	if leftIdx >= rightIdx {
		return
	}

	pivotIdx := partition(arr, leftIdx, rightIdx)
	quickSort(arr, leftIdx, pivotIdx-1)
	quickSort(arr, pivotIdx+1, rightIdx)
}

func partition(arr []*numWithIdx, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx+1)
	arr[pivotIdx], arr[rightIdx] = arr[rightIdx], arr[pivotIdx]

	swapIdx := leftIdx
	for idx := leftIdx; idx < rightIdx; idx += 1 {
		if arr[idx].Num < arr[rightIdx].Num {
			arr[swapIdx], arr[idx] = arr[idx], arr[swapIdx]
			swapIdx += 1
		}
	}

	arr[swapIdx], arr[rightIdx] = arr[rightIdx], arr[swapIdx]

	return swapIdx
}

// @lc code=end

