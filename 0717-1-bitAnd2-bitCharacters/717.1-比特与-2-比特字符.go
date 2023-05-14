/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1 比特与 2 比特字符
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {
	// solution1: 正序遍历，确保遍历到的末尾索引与len(bits)-1是否相等
	/**
	idx := 0
	for idx < len(bits)-1 {
		if bits[idx] == 0 {
			idx += 1
		} else {
			idx += 2
		}
	}

	return idx == len(bits)-1
	*/

	// solution2: 逆序，题目假设比特数组是合法的，那最后一定是10或者0
	//			  只需要遍历倒数第2个0之后1的数量即可，偶数则为ture
	idx := len(bits) - 2
	for idx >= 0 && bits[idx] == 1 {
		idx -= 1
	}

	return (len(bits)-2-idx)%2 == 0
}

// @lc code=end

