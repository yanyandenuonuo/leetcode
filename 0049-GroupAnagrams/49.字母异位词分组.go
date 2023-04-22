/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// solution1: 对每个字符进行排序，然后利用map进行聚合
	// solution2: 利用数组占位符作为map的key，直接比较占位符进行聚合
	bitCharMap := make(map[[26]int][]string)

	for _, str := range strs {
		bitChar := [26]int{}
		for _, char := range str {
			bitChar[char-'a'] += 1
		}
		if _, isExist := bitCharMap[bitChar]; !isExist {
			bitCharMap[bitChar] = make([]string, 0)
		}
		bitCharMap[bitChar] = append(bitCharMap[bitChar], str)
	}
	res := make([][]string, 0, len(bitCharMap))
	for _, strList := range bitCharMap {
		res = append(res, strList)
	}
	return res
}

// @lc code=end

