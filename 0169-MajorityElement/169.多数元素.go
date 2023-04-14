/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	frequenceCount := 0
	frequenceNum := 0
	for _, val := range nums {
		if frequenceCount <= 0 {
			frequenceNum = val
			frequenceCount = 1
		} else if frequenceNum == val {
			frequenceCount += 1
		} else if frequenceNum != val {
			frequenceCount -= 1
		}
	}
	return frequenceNum
}

// @lc code=end

