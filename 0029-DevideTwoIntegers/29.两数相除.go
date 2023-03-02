/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 {
		if divisor == -1 {
			return math.MaxInt32
		} else if divisor == 1 {
			return math.MinInt32
		}
	} else if dividend == math.MaxInt32 {
		if divisor == -1 {
			return -math.MaxInt32
		} else if divisor == 1 {
			return math.MaxInt32
		}
	}

	reverse := false
	if dividend > 0 {
		reverse = !reverse
		dividend = -dividend
	}
	if divisor > 0 {
		reverse = !reverse
		divisor = -divisor
	}

	res := 0
	for minIdx, maxIdx := 1, math.MaxInt32; minIdx <= maxIdx; {
		pivotIdx := minIdx + (maxIdx-minIdx)/2
		switch compareMultiplication(divisor, pivotIdx, dividend) {
		case 0:
			res = pivotIdx
			minIdx = math.MaxInt32
		case 1:
			minIdx = pivotIdx + 1
			res = pivotIdx
		case -1:
			maxIdx = pivotIdx - 1
		}
	}

	if reverse {
		return -res
	}

	return res
}

func compareMultiplication(a, b, c int) int {
	ans := 0
	for b > 0 {
		if b&0x01 == 0x01 {
			if c-a > ans {
				return -1
			}
			ans += a
		}
		b >>= 1
		a += a
	}

	if ans > c {
		return 1
	} else if ans < c {
		return -1
	}

	return 0
}

// @lc code=end

