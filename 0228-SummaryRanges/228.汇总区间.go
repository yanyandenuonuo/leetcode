/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	res := make([]string, 0, len(nums))
	for leftIdx := 0; leftIdx < len(nums); {
		rightIdx := leftIdx + 1
		for ; rightIdx < len(nums); rightIdx += 1 {
			if nums[rightIdx]-nums[rightIdx-1] != 1 {
				break
			}
		}

		rightIdx -= 1
		if leftIdx == rightIdx {
			res = append(res, strconv.Itoa(nums[leftIdx]))
		} else {
			res = append(res, strconv.Itoa(nums[leftIdx])+"->"+strconv.Itoa(nums[rightIdx]))
		}
		leftIdx = rightIdx + 1
	}
	return res
}

// @lc code=end

