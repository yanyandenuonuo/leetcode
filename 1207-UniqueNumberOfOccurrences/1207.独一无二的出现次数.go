/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	counter_map := make(map[int]int)
	for _, num := range arr {
		counter_map[num] += 1
	}

	frequence_map := make(map[int]int)
	for _, counter := range counter_map {
		frequence_map[counter] += 1

		if frequence_map[counter] > 1 {
			return false
		}
	}

	return true
}

// @lc code=end

