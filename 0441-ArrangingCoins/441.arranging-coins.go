/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 *
 * https://leetcode.com/problems/arranging-coins/description/
 *
 * algorithms
 * Easy (42.10%)
 * Likes:    739
 * Dislikes: 736
 * Total Accepted:    174K
 * Total Submissions: 412.6K
 * Testcase Example:  '5'
 *
 * You have a total of n coins that you want to form in a staircase shape,
 * where every k-th row must have exactly k coins.
 * ⁠
 * Given n, find the total number of full staircase rows that can be formed.
 *
 * n is a non-negative integer and fits within the range of a 32-bit signed
 * integer.
 *
 * Example 1:
 *
 * n = 5
 *
 * The coins can form the following rows:
 * ¤
 * ¤ ¤
 * ¤ ¤
 *
 * Because the 3rd row is incomplete, we return 2.
 *
 *
 *
 * Example 2:
 *
 * n = 8
 *
 * The coins can form the following rows:
 * ¤
 * ¤ ¤
 * ¤ ¤ ¤
 * ¤ ¤
 *
 * Because the 4th row is incomplete, we return 3.
 *
 *
 */

// @lc code=start
func arrangeCoins(n int) int {
	leftIdx := 0
	for rightIdx := (n + 1) / 2; leftIdx <= rightIdx; {
		midNum := leftIdx + (rightIdx-leftIdx)/2
		sum := midNum * (midNum + 1) / 2
		if sum > n {
			rightIdx = midNum - 1
		} else if sum < n {
			leftIdx = midNum + 1
		} else {
			return midNum
		}
	}
	return leftIdx - 1
}

// @lc code=end

