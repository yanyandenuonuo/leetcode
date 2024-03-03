/*
 * @lc app=leetcode.cn id=2456 lang=golang
 *
 * [2456] 最流行的视频创作者
 */

// @lc code=start
func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	res := make([][]string, 0, len(creators))
	creatorViewMap := make(map[string]int64, len(creators))
	creatorMaxViewMap := make(map[string]int, len(creators))
	creatorMinIdMap := make(map[string]string, len(creators))

	maxViewCount, creatorMap := int64(0), make(map[string]bool, len(creators))
	for idx, creator := range creators {
		creatorViewMap[creator] += int64(views[idx])

		if creatorViewMap[creator] > maxViewCount {
			maxViewCount = creatorViewMap[creator]
			creatorMap = map[string]bool{creator: true}
		} else if creatorViewMap[creator] == maxViewCount {
			creatorMap[creator] = true
		}

		if views[idx] > creatorMaxViewMap[creator] {
			creatorMaxViewMap[creator] = views[idx]
			creatorMinIdMap[creator] = ids[idx]
		} else if views[idx] == creatorMaxViewMap[creator] && (ids[idx] < creatorMinIdMap[creator] || creatorMinIdMap[creator] == "") {
			creatorMinIdMap[creator] = ids[idx]
		}
	}

	for creator := range creatorMap {
		res = append(res, []string{creator, creatorMinIdMap[creator]})
	}

	return res
}

// @lc code=end

