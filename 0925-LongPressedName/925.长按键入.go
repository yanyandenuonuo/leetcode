/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	nameIdx, typedIdx := 0, 0
	for typedIdx < len(typed) {
		if nameIdx < len(name) && name[nameIdx] == typed[typedIdx] {
			nameIdx += 1
			typedIdx += 1
			continue
		}

		if typedIdx == 0 || typed[typedIdx-1] != typed[typedIdx] {
			return false
		}

		typedIdx += 1
	}

	return nameIdx == len(name)
}

// @lc code=end

