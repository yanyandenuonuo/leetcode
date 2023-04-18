/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	// solution1: 递归
	/**
	if target <= 0 {
		return [][]int{}
	}

	if len(candidates) == 0 {
		return [][]int{}
	} else {
		res := make([][]int, 0, len(candidates))
		for idx, val := range candidates {
			resList := make([][]int, 0, len(candidates))
			if val == target {
				resList = append(resList, []int{val})
			}
			for _, subRes := range combinationSum(candidates[idx:], target-val) {
				if len(subRes) == 0 {
					continue
				}
				resList = append(resList, append([]int{val}, subRes...))
			}
			res = append(res, resList...)
		}
		return res
	}
	*/

	// solution2: DFS
	res := make([][]int, 0, len(candidates))

	dfs(candidates, 0, target, make([]int, 0, len(candidates)), &res)
	return res
}

func dfs(candidates []int, idx, target int, comb []int, res *[][]int) {
	if idx == len(candidates) {
		return
	}

	if target == 0 {
		*res = append(*res, append([]int{}, comb...))
		return
	}

	if target-candidates[idx] >= 0 {
		comb = append(comb, candidates[idx])
		dfs(candidates, idx, target-candidates[idx], comb, res)
		comb = comb[:len(comb)-1]
	}

	dfs(candidates, idx+1, target, comb, res)
}

// @lc code=end

