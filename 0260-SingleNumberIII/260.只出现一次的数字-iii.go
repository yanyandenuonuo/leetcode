/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	// solution1: hash计数
	// solution2: 位运算，将数组依次进行异或操作后可以得到z，则z = x^y
	//			  取z的最低位1(z & -z)做为区分，则必有x的该位为1或y的该位为1
	//			  利用该位将该数组分为2类，分别进行异或操作则可分别获得x，y
	z := 0
	for _, num := range nums {
		z ^= num
	}

	x, y := 0, 0

	bitZ := z // 获取最低位的1
	if z != -1<<31 {
		// 防止溢出
		bitZ = z & -z
	}

	for _, num := range nums {
		if num&bitZ == bitZ {
			x ^= num
		} else {
			y ^= num
		}
	}

	// return []int{x, x^z}

	return []int{x, y}
}

// @lc code=end

