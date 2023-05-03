/*
 * @lc app=leetcode.cn id=792 lang=golang
 *
 * [792] 匹配子序列的单词数
 */

// @lc code=start
func numMatchingSubseq(s string, words []string) int {
	// solution1: 逐个匹配
	// solution2: 逐个匹配，优化匹配过程，记录字符串每个字符的索引位置，这样匹配时可以更高效
	sIdxList := make([][]int, 26)
	for idx, runeChar := range s {
		sIdxList[runeChar-'a'] = append(sIdxList[runeChar-'a'], idx)
	}

	counter := 0
	for _, word := range words {
		if isSubSeq(word, s, sIdxList) {
			counter += 1
		}
	}

	return counter

	// solution3: 多指针指向words各个单词，遍历字符串，反向匹配words中的单词并移动符合条件的指针
	// todo
}

func isSubSeq(word, s string, sIdxList [][]int) bool {
	if len(word) > len(s) {
		return false
	}

	wordIdx := 0
	for sIdx := 0; wordIdx < len(word) && sIdx < len(s); wordIdx += 1 {
		nextIdx := findNextIdx(sIdxList[word[wordIdx]-'a'], sIdx)

		if nextIdx == -1 {
			return false
		}

		sIdx = nextIdx
	}

	return wordIdx == len(word)
}

// findNextIdx 找到大于等于val的下一个值
func findNextIdx(nums []int, val int) int {
	if len(nums) == 0 {
		return -1
	}

	leftIdx, rightIdx := 0, len(nums)-1

	if nums[leftIdx] > val {
		return nums[leftIdx] + 1
	}

	if nums[rightIdx] < val {
		return -1
	}

	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if nums[midIdx] < val {
			leftIdx = midIdx + 1
		} else if nums[midIdx] > val {
			rightIdx = midIdx - 1
		} else {
			return nums[midIdx] + 1
		}
	}

	return nums[leftIdx] + 1
}

// @lc code=end

