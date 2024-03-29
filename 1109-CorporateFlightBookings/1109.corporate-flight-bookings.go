/*
 * @lc app=leetcode id=1109 lang=golang
 *
 * [1109] Corporate Flight Bookings
 *
 * https://leetcode.com/problems/corporate-flight-bookings/description/
 *
 * algorithms
 * Medium (53.61%)
 * Likes:    544
 * Dislikes: 103
 * Total Accepted:    20.4K
 * Total Submissions: 38.1K
 * Testcase Example:  '[[1,2,10],[2,3,20],[2,5,25]]\n5'
 *
 * There are n flights, and they are labeled from 1 to n.
 *
 * We have a list of flight bookings.  The i-th booking bookings[i] = [i, j, k]
 * means that we booked k seats from flights labeled i to j inclusive.
 *
 * Return an array answer of length n, representing the number of seats booked
 * on each flight in order of their label.
 *
 *
 * Example 1:
 *
 *
 * Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
 * Output: [10,55,45,25,25]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= bookings.length <= 20000
 * 1 <= bookings[i][0] <= bookings[i][1] <= n <= 20000
 * 1 <= bookings[i][2] <= 10000
 *
 */

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	for _, val := range bookings {
		res[val[0]-1] += val[2]
		if val[1] < n {
			res[val[1]] -= val[2]
		}
	}
	for idx := 1; idx < n; idx += 1 {
		res[idx] += res[idx-1]
	}

	return res
}

// @lc code=end

