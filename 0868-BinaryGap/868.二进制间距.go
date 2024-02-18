/*
 * @lc app=leetcode.cn id=868 lang=golang
 *
 * [868] 二进制间距
 */

// @lc code=start
func binaryGap(n int) int {
	maxGap := 0
	preBitIdx := 63

	for bitIdx := 0; n > 0; bitIdx += 1 {
		if n&0x01 != 0x01 {
			n >>= 1
			continue
		}

		if bitIdx-preBitIdx > maxGap {
			maxGap = bitIdx - preBitIdx
		}

		preBitIdx = bitIdx
		n >>= 1
	}

	return maxGap
}

// @lc code=end

