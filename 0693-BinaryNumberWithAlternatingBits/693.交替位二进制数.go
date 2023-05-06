/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] 交替位二进制数
 */

// @lc code=start
func hasAlternatingBits(n int) bool {
	// solution1: 模拟
	// solution2: 位运算，采用反证法
	//			  如果n为交替位二进制，则右移一位再与n异或后的位全部为1(不考虑前导0)
	//			  得到的二进制全为1的数+1后则发生进位，低位全部为0，最高位向左进一位为1，再和其自身进行与位运算则结果一定为0
	onlyOneBit := n ^ (n >> 1)
	return onlyOneBit&(onlyOneBit+1) == 0
}

// @lc code=end

