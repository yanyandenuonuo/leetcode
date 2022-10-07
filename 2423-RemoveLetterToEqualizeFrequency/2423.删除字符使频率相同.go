/*
 * @lc app=leetcode.cn id=2423 lang=golang
 *
 * [2423] 删除字符使频率相同
 */

// @lc code=start
func equalFrequency(word string) bool {
	if len(word) == 0 {
		return true
	}

	runeMap := make(map[rune]int, len(word))
	for _, runeChar := range word {
		runeMap[runeChar] += 1
	}

	frequencyMap := make(map[int]int, len(runeMap))
	for _, frequency := range runeMap {
		frequencyMap[frequency] += 1
	}

	if len(frequencyMap) == 2 {
		// one char show time more one than other char
		// one char show time only one and others show same time
		if len(word)%len(runeMap) == 1 || len(word)%(len(runeMap)-1) == 1 {
			return true
		}
	} else if len(frequencyMap) == 1 {
		frequency := 0
		for key := range frequencyMap {
			frequency = key
		}

		// all char show time is one
		// only one char show any time
		if frequency == 1 || len(runeMap) == 1 {
			return true
		}
	}

	return false
}

// @lc code=end

