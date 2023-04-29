/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	// solution1: hash计数
	// solution2: 按位统计，因为除了目标元素外所有元素都出现了3次，所以按位对数组元素求和，对3取余后即为目标元素在该位的值
	// int在32位机器中是4个字节，在64位机器中是8个字节，所以这里声明为4个字节以排除对负数的影响
	res := int32(0)
	for idx := 0; idx < 32; idx += 1 {
		bitSum := 0
		for _, num := range nums {
			bitSum += num >> idx & 0x01
		}
		if bitSum%3 > 0 {
			res |= 0x01 << idx
		}
	}
	return int(res)

	// solution3: 利用数字电路的方式（与、或、非、异或）
	// todo
}

// @lc code=end

