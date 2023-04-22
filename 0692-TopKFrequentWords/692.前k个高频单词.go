/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
type wordFrequent struct {
	frequent int
	word     string
}

func topKFrequent(words []string, k int) []string {
	frequentMap := make(map[string]wordFrequent, len(words))
	for _, word := range words {
		if _, isExist := frequentMap[word]; !isExist {
			frequentMap[word] = wordFrequent{
				frequent: 1,
				word:     word,
			}
			continue
		}
		wordFreq := frequentMap[word]
		wordFreq.frequent += 1
		frequentMap[word] = wordFreq
	}

	frequentList := make([]wordFrequent, 0, k)
	for _, wordFreq := range frequentMap {
		if len(frequentList) < k {
			frequentList = append(frequentList, wordFrequent{
				frequent: wordFreq.frequent,
				word:     wordFreq.word,
			})
		} else {
			if wordFreq.frequent > frequentList[0].frequent ||
				(wordFreq.frequent == frequentList[0].frequent && wordFreq.word < frequentList[0].word) {
				frequentList[0] = wordFrequent{
					frequent: wordFreq.frequent,
					word:     wordFreq.word,
				}
			}
		}

		if len(frequentList) == k {
			heapSort(frequentList)
		}
	}
	res := make([]string, 0, k)
	for idx := len(frequentList) - 1; idx >= 0; idx -= 1 {
		res = append(res, frequentList[idx].word)
	}
	return res
}

func heapSort(nums []wordFrequent) {
	for idx := len(nums)/2 - 1; idx >= 0; idx -= 1 {
		heapify(nums, idx, len(nums))
	}
	for idx := len(nums) - 1; idx > 0; idx -= 1 {
		nums[0], nums[idx] = nums[idx], nums[0]
		heapify(nums, 0, idx)
	}
}

func heapify(nums []wordFrequent, idx, length int) {
	leftIdx := idx*2 + 1
	rightIdx := idx*2 + 2
	largestIdx := idx

	if leftIdx < length && (nums[leftIdx].frequent > nums[largestIdx].frequent ||
		(nums[leftIdx].frequent == nums[largestIdx].frequent && nums[leftIdx].word < nums[largestIdx].word)) {
		largestIdx = leftIdx
	}

	if rightIdx < length && (nums[rightIdx].frequent > nums[largestIdx].frequent ||
		(nums[rightIdx].frequent == nums[largestIdx].frequent && nums[rightIdx].word < nums[largestIdx].word)) {
		largestIdx = rightIdx
	}
	if largestIdx != idx {
		nums[largestIdx], nums[idx] = nums[idx], nums[largestIdx]
		heapify(nums, largestIdx, length)
	}
}

// @lc code=end

