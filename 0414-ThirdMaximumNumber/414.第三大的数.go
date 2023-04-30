/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	numRank := [3]int64{
		-1 << 63,
		-1 << 63,
		-1 << 63,
	}
	for _, num := range nums {
		if num > int(numRank[0]) {
			numRank[0], numRank[1], numRank[2] = int64(num), numRank[0], numRank[1]
		} else if int(numRank[0]) > num && num > int(numRank[1]) {
			numRank[1], numRank[2] = int64(num), numRank[1]
		} else if int(numRank[1]) > num && num > int(numRank[2]) {
			numRank[2] = int64(num)
		}
	}
	if numRank[2] == -1<<63 {
		return int(numRank[0])
	}
	return int(numRank[2])
}

// @lc code=end

