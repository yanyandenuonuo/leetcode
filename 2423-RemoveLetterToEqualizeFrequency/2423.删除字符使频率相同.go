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
		if len(frequencyMap) > 2 {
			return false
		}
	}

	if len(frequencyMap) == 2 {
		minFrequency := len(word)
		maxFrequency := 0
		minFrequencyCount := 0
		maxFrequencyCount := 0
		for frequency, frequencyCount := range frequencyMap {
			if frequency < minFrequency {
				minFrequency = frequency
				minFrequencyCount = frequencyCount
			}
			if frequency > maxFrequency {
				maxFrequency = frequency
				maxFrequencyCount = frequencyCount
			}
		}

		// one char show time more one than other char
		// one char show time only one and others show same time
		if (maxFrequency-minFrequency == 1 && maxFrequencyCount == 1) ||
			(minFrequency == 1 && minFrequencyCount == 1) {
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

