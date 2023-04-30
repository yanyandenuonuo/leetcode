/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	// solution1: 逐位操作
	/**
	reverseNum := uint32(0)
	for idx := 0; idx < 32; idx += 1 {
		reverseNum <<= 1
		reverseNum |= num & 0x01
		num >>= 1
	}
	return reverseNum
	*/

	// solution2: 逐次批量交换位值，本质上是递归翻转的思路，先翻转1bit，再翻转2bit，4bit，8bit，16bit
	//			  通过与或运算实现对应bit的翻转
	//			  1bit	1 2 3 4 5 6 7 8
	//			  2bit	2 1 4 3 6 5 8 7
	//			  4bit	4 3 2 1 8 7 6 5
	//			  8bit	8 7 6 5 4 3 2 1
	const (
		oneBit   = 0b01010101010101010101010101010101
		twoBit   = 0b00110011001100110011001100110011
		fourBit  = 0b00001111000011110000111100001111
		eightBit = 0b00000000111111110000000011111111
	)
	num = num>>1&oneBit | num&oneBit<<1
	num = num>>2&twoBit | num&twoBit<<2
	num = num>>4&fourBit | num&fourBit<<4
	num = num>>8&eightBit | num&eightBit<<8
	return num>>16 | num<<16
}

// @lc code=end

