/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	// solution1: 暴力
	// solution2: 暴力优化版，使用一个数组记录温度出现的最早时间，通过倒序遍历的方式逐次求解
	/**
	res := make([]int, len(temperatures))
	idxTemperatures := [71]int{} // 温度值范围在30-100，这里做了偏移处理，针对idxTemperatures的查询全部-30以节省存储空间

	for idx := len(temperatures) - 1; idx >= 0; idx -= 1 {
		idxTemperatures[temperatures[idx]-30] = idx

		for idxTemperature := temperatures[idx] - 30 + 1; idxTemperature <= 70; idxTemperature += 1 {
			if idxTemperatures[idxTemperature] == 0 {
				continue
			}
			if res[idx] == 0 || idxTemperatures[idxTemperature]-idx < res[idx] {
				res[idx] = idxTemperatures[idxTemperature] - idx
			}
		}
	}
	return res
	*/

	// solution3: 利用栈空间，使用栈存放未找到答案的温度值的索引值，每个元素和栈顶元素做对比，若符合则出栈，直到栈为空或不符合条件，将元素放入
	res := make([]int, len(temperatures))
	stackTemperature := make([]int, 0, len(temperatures))
	for idx := range temperatures {
		for len(stackTemperature) != 0 && temperatures[idx] > temperatures[stackTemperature[len(stackTemperature)-1]] {
			res[stackTemperature[len(stackTemperature)-1]] = idx - stackTemperature[len(stackTemperature)-1]
			stackTemperature = stackTemperature[:len(stackTemperature)-1]
		}
		stackTemperature = append(stackTemperature, idx)
	}

	return res
}

// @lc code=end

