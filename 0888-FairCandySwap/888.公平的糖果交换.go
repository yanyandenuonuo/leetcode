/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果交换
 */

// @lc code=start
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceTotal := 0
	bobTotal := 0

	for _, aliceCandy := range aliceSizes {
		aliceTotal += aliceCandy
	}

	for _, bobCandy := range bobSizes {
		bobTotal += bobCandy
	}

	if (aliceTotal+bobTotal)&0x01 == 0x01 {
		return []int{}
	}

	diff := aliceTotal - (aliceTotal+bobTotal)/2

	for _, aliceCandy := range aliceSizes {
		for _, bobCandy := range bobSizes {
			if aliceCandy-bobCandy == diff {
				return []int{aliceCandy, bobCandy}
			}
		}
	}

	return []int{}
}

// @lc code=end

