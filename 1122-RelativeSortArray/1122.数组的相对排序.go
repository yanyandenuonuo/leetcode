/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	minValue := 1000
	maxValue := 0
	for _, val := range arr1 {
		if val > maxValue {
			maxValue = val
		}
		if val < minValue {
			minValue = val
		}
	}

	frequencyCounter := make([]int, maxValue-minValue+1)
	for _, val := range arr1 {
		idx := val - minValue
		frequencyCounter[idx] += 1
	}
	res := make([]int, 0, len(arr1))
	for _, val := range arr2 {
		idx := val - minValue
		for frequencyCounter[idx] > 0 {
			res = append(res, val)
			frequencyCounter[idx] -= 1
		}
	}
	for idx, frequency := range frequencyCounter {
		val := minValue + idx
		for frequency > 0 {
			res = append(res, val)
			frequency -= 1
		}
	}
	return res
}

// @lc code=end

