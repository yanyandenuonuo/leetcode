/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */

// @lc code=start
type scoreWithIdx struct {
	val   int
	index int
}

func findRelativeRanks(score []int) []string {
	scoreList := make([]scoreWithIdx, 0, len(score))
	for idx := range score {
		scoreList = append(scoreList, scoreWithIdx{
			val:   score[idx],
			index: idx,
		})
	}
	quickSort(scoreList, 0, len(scoreList)-1)
	res := make([]string, len(score))
	for idx, scoreIdx := range scoreList {
		switch idx {
		case 0:
			res[scoreIdx.index] = "Gold Medal"
		case 1:
			res[scoreIdx.index] = "Silver Medal"
		case 2:
			res[scoreIdx.index] = "Bronze Medal"
		default:
			res[scoreIdx.index] = strconv.Itoa(idx + 1)
		}
	}
	return res
}

func quickSort(nums []scoreWithIdx, leftIdx, rightIdx int) {
	if leftIdx >= rightIdx {
		return
	}
	pivotIdx := partition(nums, leftIdx, rightIdx)
	quickSort(nums, leftIdx, pivotIdx-1)
	quickSort(nums, pivotIdx+1, rightIdx)
}

func partition(nums []scoreWithIdx, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx+1)
	nums[pivotIdx], nums[rightIdx] = nums[rightIdx], nums[pivotIdx]
	swapIdx := leftIdx
	for idx := leftIdx; idx < rightIdx; idx += 1 {
		if nums[idx].val > nums[rightIdx].val {
			nums[idx], nums[swapIdx] = nums[swapIdx], nums[idx]
			swapIdx += 1
		}
	}
	nums[swapIdx], nums[rightIdx] = nums[rightIdx], nums[swapIdx]
	return swapIdx
}

// @lc code=end

