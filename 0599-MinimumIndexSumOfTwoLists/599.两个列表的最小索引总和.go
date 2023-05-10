/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	// solution: hash
	restaurantMap := make(map[string]int, len(list1))
	for idx, restaurant := range list1 {
		restaurantMap[restaurant] = idx
	}

	resList, minSum := make([]string, 0), len(list1)+len(list2)

	for idx, restaurant := range list2 {
		if _, isExist := restaurantMap[restaurant]; !isExist {
			continue
		}

		if restaurantMap[restaurant]+idx < minSum {
			resList = []string{restaurant}
			minSum = restaurantMap[restaurant] + idx
		} else if restaurantMap[restaurant]+idx == minSum {
			resList = append(resList, restaurant)
		}
	}

	return resList
}

// @lc code=end

