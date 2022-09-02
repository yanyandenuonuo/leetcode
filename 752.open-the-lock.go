/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 *
 * https://leetcode.com/problems/open-the-lock/description/
 *
 * algorithms
 * Medium (52.08%)
 * Likes:    1207
 * Dislikes: 46
 * Total Accepted:    71.5K
 * Total Submissions: 137.3K
 * Testcase Example:  '["0201","0101","0102","1212","2002"]\n"0202"'
 *
 * You have a lock in front of you with 4 circular wheels. Each wheel has 10
 * slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can
 * rotate freely and wrap around: for example we can turn '9' to be '0', or '0'
 * to be '9'. Each move consists of turning one wheel one slot.
 * 
 * The lock initially starts at '0000', a string representing the state of the
 * 4 wheels.
 * 
 * You are given a list of deadends dead ends, meaning if the lock displays any
 * of these codes, the wheels of the lock will stop turning and you will be
 * unable to open it.
 * 
 * Given a target representing the value of the wheels that will unlock the
 * lock, return the minimum total number of turns required to open the lock, or
 * -1 if it is impossible.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * Output: 6
 * Explanation:
 * A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" ->
 * "1201" -> "1202" -> "0202".
 * Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202"
 * would be invalid,
 * because the wheels of the lock become stuck after the display becomes the
 * dead end "0102".
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: deadends = ["8888"], target = "0009"
 * Output: 1
 * Explanation:
 * We can turn the last wheel in reverse to move from "0000" -> "0009".
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
 * target = "8888"
 * Output: -1
 * Explanation:
 * We can't reach the target without getting stuck.
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: deadends = ["0000"], target = "8888"
 * Output: -1
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= deadends.length <= 500
 * deadends[i].length == 4
 * target.length == 4
 * target will not be in the list deadends.
 * target and deadends[i] consist of digits only.
 * 
 * 
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	if len(target) > 4 {
		return -1
	}

	startKey := "0000"
	if startKey == target {
		return 0
	}
	
	deadMap := make(map[string]bool, len(deadends))
	for _, deadKey := range deadends {
		deadMap[deadKey] = true
	}

	handleMap := make(map[string]bool)

	queueList := make([]string, 0, 64)
	queueList = append(queueList, startKey)
	handleMap[startKey] = true
	step := 0
	for len(queueList) > 0 {
		currentQueue := len(queueList)
		for idx := 0; idx < currentQueue; idx += 1 {
			currentKey := queueList[idx]


			if currentKey == target {
				return step
			}

			if _, isDeadKey := deadMap[currentKey]; isDeadKey {
				continue
			}

			for bitIdx := 0; bitIdx < 4; bitIdx += 1 {
				currentBitNum, _ := strconv.Atoi(currentKey[bitIdx : bitIdx+1])
				nextBitNum := (currentBitNum + 1) % 10
				nextKey := ""
				if bitIdx == 3 {
					nextKey = currentKey[0:bitIdx] + strconv.Itoa(nextBitNum)
				} else {
					nextKey = currentKey[0:bitIdx] + strconv.Itoa(nextBitNum) + currentKey[bitIdx+1:]
				}
				if _, isHandle := handleMap[nextKey]; !isHandle {
					queueList = append(queueList, nextKey)
					handleMap[nextKey] = true
				}

				prevBitNum := (currentBitNum + 10 - 1) % 10
				prevKey := ""
				if bitIdx == 3 {
					prevKey = currentKey[0:bitIdx] + strconv.Itoa(prevBitNum)
				} else {
					prevKey = currentKey[0:bitIdx] + strconv.Itoa(prevBitNum) + currentKey[bitIdx+1:]
				}
				if _, isHandle := handleMap[prevKey]; !isHandle {
					queueList = append(queueList, prevKey)
					handleMap[prevKey] = true
				}
			}
		}

		if len(queueList) > currentQueue {
			queueList = queueList[currentQueue:]
		} else {
			queueList = make([]string, 0)
		}
		step += 1
	}

	return -1
}
// @lc code=end

