/*
 * @lc app=leetcode.cn id=1615 lang=golang
 *
 * [1615] 最大网络秩
 */

// @lc code=start
func maximalNetworkRank(n int, roads [][]int) int {
	cityCount := make(map[int]int, n)
	cityConnect := make(map[int]map[int]int, n)
	cityList := make([]int, 0, n)

	for idx := 0; idx < len(roads); idx += 1 {
		cityA := roads[idx][0]
		cityB := roads[idx][1]

		if val, isExist := cityConnect[cityA][cityB]; isExist && val == 1 {
			// 同时存在roads[cityA][cityB]和roads[cityB][cityA]时处理一次即可
			continue
		}

		if _, isExist := cityCount[cityA]; !isExist {
			cityList = append(cityList, cityA)
			cityConnect[cityA] = make(map[int]int, n)
		}

		if _, isExist := cityCount[cityB]; !isExist {
			cityList = append(cityList, cityB)
			cityConnect[cityB] = make(map[int]int, n)
		}

		cityCount[cityA] += 1
		cityCount[cityB] += 1
		cityConnect[cityA][cityB] = 1
		cityConnect[cityB][cityA] = 1

	}
	res := 0
	// n个城市可能有些城市不存在道路，所以这里不能使用n，需要使用cityList
	for cityAIdx := 0; cityAIdx < len(cityList); cityAIdx += 1 {
		for cityBIdx := cityAIdx + 1; cityBIdx < len(cityList); cityBIdx += 1 {
			cityA := cityList[cityAIdx]
			cityB := cityList[cityBIdx]
			// 互相之间的路-1
			currentRes := cityCount[cityA] + cityCount[cityB] - cityConnect[cityA][cityB]
			if currentRes > res {
				res = currentRes
			}
		}
	}
	return res
}

// @lc code=end

