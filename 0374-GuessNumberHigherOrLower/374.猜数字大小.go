/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	for leftIdx, rightIdx := 0, n; leftIdx <= rightIdx; {
		midNum := leftIdx + (rightIdx-leftIdx)/2
		switch guess(midNum) {
		case 1:
			leftIdx = midNum + 1
		case -1:
			rightIdx = midNum - 1
		default:
			return midNum
		}
	}
	return -1
}

// @lc code=end

