/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	// solution1: 暴力求解，超时
	/**
	isPrimeFunc := func(x int) bool {
		if x < 2 {
			return false
		}
		for idx := 2; idx*idx <= x; idx += 1 {
			if x%idx == 0 {
				return false
			}
		}
		return true
	}
	total := 0
	for idx := 2; idx < n; idx += 1 {
		if isPrimeFunc(idx) {
			total += 1
		}
	}
	return total
	*/

	// solution2: 埃氏筛
	// 将质数x的2x, 3x, 4x, ... nx标记为合数

	notPrimeDict := make([]bool, n)
	total := 0
	for idx := 2; idx < n; idx += 1 {
		if !notPrimeDict[idx] {
			total += 1
			for x := idx; idx*x < n; x += 1 {
				notPrimeDict[idx*x] = true
			}
		}
	}

	return total

	// solution3: 线性筛
	// 将质数记录起来，然后将x * prime[idx]标记为合数，当x % prime[idx] == 0时停止标记（要使得合数z被最小的质数因子标记为合数）
	// 假设
	//   y = x / prime[idx] => y > 1
	// 则存在
	//   z = x * prime[idx+1] => z = y * prime[idx] * prime[idx+1] = (y * prime[idx+1]) * prime[idx]
	// 示例如下
	// 2 => 2 * [2] 			=> 4
	// 3 => 3 * [2, 3] 			=> 6 9
	// 4 => 4 * [2, 3] 			=> 8 [12]
	// 5 => 5 * [2, 3] 			=> 10, 15
	// 6 => 6 * [2, 3, 5]		=> 12, [18], [30]
	// 7 => 7 * [2, 3, 5]		=> 14, 21, 35
	// 8 => 8 * [2, 3, 5, 7]	=> 16, [24], [40], [56]
	// 9 => 9 * [2, 3, 5, 7]	=> 18, 27，[45], [63]
	/**
	primeList := make([]int, 0, n)
	notPrimeDict := make([]bool, n)
	for idx := 2; idx < n; idx += 1 {
		if !notPrimeDict[idx] {
			primeList = append(primeList, idx)
		}
		for _, prime := range primeList {
			val := idx * prime
			if val >= n {
				break
			}
			notPrimeDict[val] = true
			if idx%prime == 0 {
				break
			}
		}
	}
	return len(primeList)
	*/
}

// @lc code=end

