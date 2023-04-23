/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	// solution1: 直接排序，高位越高越靠前
	/**
	numStrList := make([]string, 0, len(nums))
	for _, num := range nums {
		numStrList = append(numStrList, strconv.Itoa(num))
	}
	quickSort(numStrList, 0, len(numStrList)-1)
	res := ""
	for _, numStr := range numStrList {
		res += numStr
	}
	if res[0] == '0' {
		res = "0"
	}
	return res
	*/

	// solution2: 先按最高位排，同序的再排，依然是高位越高越靠前
	if len(nums) == 0 {
		return ""
	} else if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	numStrList := [10][]string{}
	for _, num := range nums {
		numStr := strconv.Itoa(num)
		maxBit := numStr[0] - '0'
		numStrList[maxBit] = append(numStrList[maxBit], numStr)
	}
	res := ""
	for idx := 9; idx >= 0; idx -= 1 {
		if len(numStrList[idx]) == 0 {
			continue
		} else if len(numStrList[idx]) == 1 {
			res += numStrList[idx][0]
			continue
		} else {
			quickSort(numStrList[idx], 0, len(numStrList[idx])-1)
			for _, numStr := range numStrList[idx] {
				res += numStr
			}
		}
	}

	// fix "00000" => "0"
	if res[0] == '0' {
		res = "0"
	}

	return res
}

func quickSort(nums []string, leftIdx, rightIdx int) {
	if leftIdx >= rightIdx {
		return
	}
	pivotIdx := partitionNums(nums, leftIdx, rightIdx)
	quickSort(nums, leftIdx, pivotIdx-1)
	quickSort(nums, pivotIdx+1, rightIdx)
}

func partitionNums(nums []string, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx+1)
	nums[pivotIdx], nums[rightIdx] = nums[rightIdx], nums[pivotIdx]
	swapIdx := leftIdx
	for idx := leftIdx; idx < rightIdx; idx += 1 {
		if isGreaterOrEqual(nums[idx], nums[rightIdx]) {
			nums[idx], nums[swapIdx] = nums[swapIdx], nums[idx]
			swapIdx += 1
		}
	}
	nums[swapIdx], nums[rightIdx] = nums[rightIdx], nums[swapIdx]
	return swapIdx
}

func isGreaterOrEqual(numStr1, numStr2 string) bool {
	numStr1, numStr2 = numStr1+numStr2, numStr2+numStr1

	return !(numStr1 < numStr2)
}

// @lc code=end

