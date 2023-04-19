/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	// solution1: 枚举，基于二进制表示对应位是否选中
	/**
	res := make([][]int, 0, 1<<len(nums))
	res = append(res, []int{})

	for bitIdx := 0; bitIdx < 1<<len(nums); bitIdx += 1 {
		subset := make([]int, 0, len(nums))
		for idx, val := range nums {
			if bitIdx>>idx&0x01 == 0x01 {
				subset = append(subset, val)
			}
		}
		res = append(res, subset)
	}

	return res
	*/

	// solution2: DFS
	res := make([][]int, 0, 1<<len(nums))

	subset := make([]int, 0, len(nums))

	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, subset...))
			return
		}

		subset = append(subset, nums[idx])
		dfs(idx + 1)

		subset = subset[:len(subset)-1]
		dfs(idx + 1)
	}

	dfs(0)

	return res
}

// @lc code=end

