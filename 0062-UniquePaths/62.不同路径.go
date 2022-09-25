/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// solution1: f(m, n) = f(m-1, n) + f(m, n-1)， 注意需要将边界初始化为1
	// res := make([][]int, m)

	// for idxM := 0; idxM < m; idxM += 1 {
	// 	res[idxM] = make([]int, n)
	// 	res[idxM][0] = 1
	// }

	// for idxN := 0; idxN < n; idxN += 1 {
	// 	res[0][idxN] = 1
	// }

	// for idxM := 1; idxM < m; idxM += 1 {
	// 	for idxN := 1; idxN < n; idxN += 1 {
	// 		res[idxM][idxN] += res[idxM-1][idxN] + res[idxM][idxN-1]
	// 	}
	// }
	// return res[m-1][n-1]

	// solution2: 排列组合数，实际需要向下走(m-1)步，向右走(n-1)步，共需要走(m+n-2)步，从中选出向下走(m-1)布或者向右走(n-1)步的组合数即可。
	//  n
	// Cm = m * (m-1) * (m-2)*...*(m-n+1)/(n*(n-1)*(n-2)*...*1)
	//  n    (m-n)
	// Cm == Cm
	m = m + n - 2
	n = n - 1
	if (m - n) < n {
		n = m - n
	}

	nSum, mSum := 1, 1
	for idx := 0; idx < n; idx += 1 {
		mSum *= (m - idx)
		nSum *= (n - idx)

		if (mSum % nSum) == 0 {
			// 解决溢出问题
			mSum = mSum / nSum
			nSum = 1
		}

		// 如果没有溢出问题，则可以直接使用
		// total := 1
		// total *= (m - idx)/(n - idx)
		// 最后直接返回total即可
	}
	return mSum / nSum
}

// @lc code=end

