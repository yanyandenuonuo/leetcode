/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	// solution: 将g和s分别排序，然后依次找出符合条件的进行分配即可
	sort.Ints(g)
	sort.Ints(s)
	gIdx := 0
	for sIdx := 0; gIdx < len(g) && sIdx < len(s); {
		if g[gIdx] <= s[sIdx] {
			gIdx += 1
			sIdx += 1
			continue
		}
		sIdx += 1
	}

	return gIdx
}

// @lc code=end

