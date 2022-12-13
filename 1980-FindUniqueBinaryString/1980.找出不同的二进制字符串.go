/*
 * @lc app=leetcode.cn id=1980 lang=golang
 *
 * [1980] 找出不同的二进制字符串
 */

// @lc code=start
func findDifferentBinaryString(nums []string) string {
	resBytes := make([]byte, len(nums))
	for idx := 0; idx < len(nums); idx += 1 {
		if nums[idx][idx] == '0' {
			resBytes[idx] = '1'
		} else {
			resBytes[idx] = '0'
		}
	}
	return string(resBytes)
}

// @lc code=end

