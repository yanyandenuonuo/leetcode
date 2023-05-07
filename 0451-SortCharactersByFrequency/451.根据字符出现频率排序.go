/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start
type charWithCounter struct {
	runeChar rune
	counter  int
}

func frequencySort(s string) string {
	charCounter := make(map[rune]int)
	for _, runeChar := range s {
		charCounter[runeChar] += 1
	}

	charList := make([]charWithCounter, 0, len(charCounter))
	for runeChar := range charCounter {
		charList = append(charList, charWithCounter{
			runeChar: runeChar,
			counter:  charCounter[runeChar],
		})
	}

	quickSort(charList, 0, len(charList)-1)

	charBytes := make([]rune, 0, len(s))

	for _, currentChar := range charList {
		for idx := 0; idx < currentChar.counter; idx += 1 {
			charBytes = append(charBytes, currentChar.runeChar)
		}
	}

	return string(charBytes)
}

func quickSort(strList []charWithCounter, leftIdx, rightIdx int) {
	if leftIdx >= rightIdx {
		return
	}

	pivotIdx := partition(strList, leftIdx, rightIdx)
	quickSort(strList, leftIdx, pivotIdx-1)
	quickSort(strList, pivotIdx+1, rightIdx)
}

func partition(strList []charWithCounter, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx+1)
	strList[pivotIdx], strList[rightIdx] = strList[rightIdx], strList[pivotIdx]

	swapIdx := leftIdx

	for idx := leftIdx; idx < rightIdx; idx += 1 {
		if strList[idx].counter > strList[rightIdx].counter ||
			(strList[idx].counter == strList[rightIdx].counter && strList[idx].runeChar > strList[rightIdx].runeChar) {
			strList[idx], strList[swapIdx] = strList[swapIdx], strList[idx]
			swapIdx += 1
		}
	}

	strList[swapIdx], strList[rightIdx] = strList[rightIdx], strList[swapIdx]

	return swapIdx
}

// @lc code=end

