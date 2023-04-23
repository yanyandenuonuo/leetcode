/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// solution1: 动态规划
	/**
	maxLIS := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for idx := 1; idx < len(nums); idx += 1 {
		for leftIdx := 0; leftIdx < idx; leftIdx += 1 {
			if nums[idx] > nums[leftIdx] && dp[leftIdx]+1 > dp[idx] {
				dp[idx] = dp[leftIdx] + 1
			}
		}
		if dp[idx] == 0 {
			dp[idx] = 1
		}

		if dp[idx] > maxLIS {
			maxLIS = dp[idx]
		}
	}
	return maxLIS
	*/

	// solution2: 贪心 + 二分查找
	//			  要使上升子序列尽可能的长，则需要让序列上升得尽可能慢，每次在上升子序列最后加上的那个数尽可能的小
	//			  使用一个数组d[l]表示长度为l的最长上升子序列的末尾元素的最小值，len表示d数组长度
	//				如果nums[idx] > d[len]则直接加入到d数组末尾，并更新len=len(d)+1
	//				如果nums[idx] <= d[len]则在d数组中二分查找，找到第一个比nums[idx]小的数d[x]，更新d[x+1]=nums[idx]
	d := make([]int, 0, len(nums))
	d = append(d, nums[0])

	for idx := 1; idx < len(nums); idx += 1 {
		if nums[idx] > d[len(d)-1] {
			d = append(d, nums[idx])
			continue
		}
		targetIdx := findLessOrEqualNum(d, 0, len(d)-1, nums[idx])
		d[targetIdx+1] = nums[idx]
	}
	return len(d)
}

func findLessOrEqualNum(nums []int, leftIdx, rightIdx, target int) int {
	if leftIdx > rightIdx {
		return rightIdx
	}

	mid := leftIdx + (rightIdx-leftIdx)/2
	if nums[mid] < target {
		return findLessOrEqualNum(nums, mid+1, rightIdx, target)
	} else if nums[mid] > target {
		return findLessOrEqualNum(nums, leftIdx, mid-1, target)
	} else {
		return mid - 1
	}
}

// @lc code=end

