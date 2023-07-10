/*
 * @lc app=leetcode.cn id=1945 lang=golang
 *
 * [1945] 字符串转化后的各位数字之和
 */

// @lc code=start
func getLucky(s string, k int) int {
	charMap := map[rune]string{
		'a': "1",
		'b': "2",
		'c': "3",
		'd': "4",
		'e': "5",
		'f': "6",
		'g': "7",
		'h': "8",
		'i': "9",
		'j': "10",
		'k': "11",
		'l': "12",
		'm': "13",
		'n': "14",
		'o': "15",
		'p': "16",
		'q': "17",
		'r': "18",
		's': "19",
		't': "20",
		'u': "21",
		'v': "22",
		'w': "23",
		'x': "24",
		'y': "25",
		'z': "26",
	}

	numStr := ""
	for _, runeChar := range s {
		numStr += charMap[runeChar]
	}

	numSum := 0
	for idx := 0; idx < k; idx += 1 {
		numSum = 0
		for _, numChar := range numStr {
			numSum += int(numChar - '0')
		}
		numStr = strconv.Itoa(numSum)
	}

	return numSum
}

// @lc code=end

