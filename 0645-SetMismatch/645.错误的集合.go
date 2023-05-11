/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */

// @lc code=start
func findErrorNums(nums []int) []int {
	// solution1: 排序
	// solution2: 位排序
	// solution3: hash
	/**
	errNum, sum, numMap := 0, 0, make(map[int]bool, len(nums)-1)

	for _, num := range nums {
		if numMap[num] {
			errNum = num
		}

		numMap[num] = true
		sum += num
	}

	return []int{errNum, len(nums)*(len(nums)+1)/2 - sum + errNum}
	*/

	// solution4: 位运算，在nums后拼接[1, ..., n]则新数组中重复数字出现3次，缺失数字出现1次，其它数字出现1次
	//			  对新数组进行异或运算，则结果为x^y
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	for num := 1; num <= len(nums); num += 1 {
		xor ^= num
	}

	lowBit := xor & (-xor) // 获取最低位的1
	numX, numY := 0, 0
	for _, num := range nums {
		if num&lowBit == 0 {
			numX ^= num
		} else {
			numY ^= num
		}
	}
	for num := 1; num <= len(nums); num += 1 {
		if num&lowBit == 0 {
			numX ^= num
		} else {
			numY ^= num
		}
	}

	for _, num := range nums {
		if num == numX {
			return []int{numX, numY}
		} else if num == numY {
			return []int{numY, numX}
		}
	}

	return []int{}
}

// @lc code=end

