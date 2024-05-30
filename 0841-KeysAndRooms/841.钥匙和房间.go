/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * [841] 钥匙和房间
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	// solution1: DFS
	// solution2: BFS
	roomMap := make(map[int]bool, len(rooms))
	roomMap[0] = true

	keyList := make([]int, 0, len(rooms))
	keyList = append(keyList, rooms[0]...)

	for idx := 0; idx < len(keyList) && len(roomMap) != len(rooms); idx += 1 {
		if roomMap[keyList[idx]] {
			continue
		}

		keyList = append(keyList, rooms[keyList[idx]]...)
		roomMap[keyList[idx]] = true
	}

	return len(roomMap) == len(rooms)
}

// @lc code=end

