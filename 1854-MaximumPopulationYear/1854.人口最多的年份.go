/*
 * @lc app=leetcode.cn id=1854 lang=golang
 *
 * [1854] 人口最多的年份
 */

// @lc code=start
func maximumPopulation(logs [][]int) int {
	// solution: 拿数组存放每年的人口数，再遍历数组
	populationList := make([]int, 101)

	for _, log := range logs {
		populationList[log[0]-1950] += 1
		populationList[log[1]-1950] -= 1
	}

	population, maxPopulation, maxYear := 0, 0, 0

	for idx, val := range populationList {
		population += val

		if population > maxPopulation {
			maxPopulation = population
			maxYear = idx
		}
	}

	return maxYear + 1950
}

// @lc code=end

