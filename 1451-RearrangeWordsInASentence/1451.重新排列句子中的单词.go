/*
 * @lc app=leetcode.cn id=1451 lang=golang
 *
 * [1451] 重新排列句子中的单词
 */

// @lc code=start
type wordWithIdx struct {
	word string
	idx  int
}

func arrangeWords(text string) string {
	wordList := strings.Split(text, " ")

	wordWithIdxList := make([]wordWithIdx, len(wordList))
	for idx, word := range wordList {
		wordWithIdxList[idx] = wordWithIdx{
			word: word,
			idx:  idx,
		}
	}

	quickSort(wordWithIdxList, 0, len(wordList)-1)

	for idx, wordWithIdx := range wordWithIdxList {
		wordList[idx] = strings.ToLower(wordWithIdx.word)
	}

	wordList[0] = string(wordList[0][0]-'a'+'A') + wordList[0][1:]

	return strings.Join(wordList, " ")
}

func quickSort(strList []wordWithIdx, leftIdx, rightIdx int) {
	if leftIdx >= rightIdx {
		return
	}

	pivotIdx := partition(strList, leftIdx, rightIdx)
	quickSort(strList, leftIdx, pivotIdx-1)
	quickSort(strList, pivotIdx+1, rightIdx)
}

func partition(strList []wordWithIdx, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx+1)
	strList[pivotIdx], strList[rightIdx] = strList[rightIdx], strList[pivotIdx]

	swapIdx := leftIdx

	for idx := leftIdx; idx < rightIdx; idx += 1 {
		if len(strList[idx].word) < len(strList[rightIdx].word) ||
			(len(strList[idx].word) == len(strList[rightIdx].word) && strList[idx].idx < strList[rightIdx].idx) {
			strList[idx], strList[swapIdx] = strList[swapIdx], strList[idx]
			swapIdx += 1
		}
	}

	strList[swapIdx], strList[rightIdx] = strList[rightIdx], strList[swapIdx]

	return swapIdx
}

// @lc code=end

