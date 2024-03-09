/*
 * @lc app=leetcode.cn id=2129 lang=golang
 *
 * [2129] 将标题首字母大写
 */

// @lc code=start
func capitalizeTitle(title string) string {
	charList := make([]string, len(title))
	for leftIdx := 0; leftIdx < len(title); {
		if title[leftIdx] == ' ' {
			charList[leftIdx] = " "
			leftIdx += 1
			continue
		}

		rightIdx := leftIdx
		for rightIdx < len(title) {
			if title[rightIdx] == ' ' {
				charList[rightIdx] = " "
				break
			} else if 'A' <= title[rightIdx] && title[rightIdx] <= 'Z' {
				charList[rightIdx] = string(title[rightIdx] - 'A' + 'a')
			} else {
				charList[rightIdx] = title[rightIdx : rightIdx+1]
			}

			rightIdx += 1
		}

		if rightIdx-leftIdx > 2 {
			charList[leftIdx] = string(charList[leftIdx][0] - 'a' + 'A')
		}

		leftIdx = rightIdx
	}

	return strings.Join(charList, "")
}

// @lc code=end

