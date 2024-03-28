/*
 * @lc app=leetcode.cn id=2333 lang=golang
 *
 * [2333] 最小差值平方和
 */

// @lc code=start
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	// solution: 将差值排序，逐次减最大的差值
	//			 将num[i]减至nums[i-1]
	//			 将num[i],nums[i-1]减至nums[i-2]
	//			 将num[i],nums[i-1],nums[i-2]减至nums[i-3]
	//			 ...
	//			 为了不判断nums[i-x]的有效性，可以补一个值为0的diff边界
	//			 如果sum(diff) <= k 可以直接返回0
	diffNum, diffSum, sumSquareDiff := make([]int, len(nums1)+1), 0, int64(0)
	k := k1 + k2
	for idx := range nums1 {
		if nums1[idx] > nums2[idx] {
			diffNum[idx+1] = nums1[idx] - nums2[idx]
		} else {
			diffNum[idx+1] = nums2[idx] - nums1[idx]
		}

		diffSum += diffNum[idx+1]

		sumSquareDiff += int64(diffNum[idx+1] * diffNum[idx+1])
	}

	if diffSum <= k {
		return 0
	}

	slices.Sort(diffNum)

	for idx := len(diffNum) - 1; idx > 0 && k > 0; idx -= 1 {
		sumSquareDiff -= int64(diffNum[idx] * diffNum[idx])

		diffSum = (diffNum[idx] - diffNum[idx-1]) * (len(diffNum) - idx)
		if diffSum < k {
			k -= diffSum
		} else if diffSum > k {
			// 把k均分到这段范围的值
			avgK := k / (len(diffNum) - idx)
			minCount := k % (len(diffNum) - idx)

			// 如果minCount>0则最小值应该再减1
			sumSquareDiff += int64(minCount * (diffNum[idx] - avgK - 1) * (diffNum[idx] - avgK - 1))

			sumSquareDiff += int64((len(diffNum) - idx - minCount) * (diffNum[idx] - avgK) * (diffNum[idx] - avgK))
			break
		} else {
			sumSquareDiff += int64((len(diffNum) - idx) * diffNum[idx-1] * diffNum[idx-1])
			break
		}
	}

	return sumSquareDiff
}

// @lc code=end

