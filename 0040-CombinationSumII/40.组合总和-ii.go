/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	// solution1: 回溯，排序+计算频率避免重复计算
	/**
	sort.Ints(candidates)

	numCounterMap := make(map[int]int, 0)
	for _, candidate := range candidates {
		numCounterMap[candidate] += 1
	}

	res := make([][]int, 0)
	numList := make([]int, 0, len(candidates))

	var dfs func(idx, targetNum int)
	dfs = func(idx, targetNum int) {
		if targetNum == 0 {
			res = append(res, append([]int{}, numList...))
			return
		}

		if idx == len(candidates) || targetNum < 0 {
			return
		}

		for offset := 0; offset < numCounterMap[candidates[idx]]; offset += 1 {
			numList = append(numList, candidates[idx])
			dfs(idx+numCounterMap[candidates[idx]], targetNum-(offset+1)*candidates[idx])
		}

		numList = numList[:len(numList)-numCounterMap[candidates[idx]]]

		dfs(idx+numCounterMap[candidates[idx]], targetNum)
	}

	dfs(0, target)

	return res
	*/

	// solution2: 回溯，计算频率+唯一数组避免重复计算
	numCounterMap := make(map[int]int, 0)
	uniqueCandidateList := make([]int, 0)

	for _, candidate := range candidates {
		if _, isExist := numCounterMap[candidate]; !isExist {
			uniqueCandidateList = append(uniqueCandidateList, candidate)
		}
		numCounterMap[candidate] += 1
	}

	var findNum func(numList []int, targetNum int) [][]int
	findNum = func(numList []int, targetNum int) [][]int {
		if len(numList) == 0 {
			return [][]int{}
		} else if len(numList) == 1 {
			if numList[0] == targetNum {
				return [][]int{{numList[0]}}
			}
		}

		resList := make([][]int, 0)

		for idx := 0; idx < len(numList); idx += 1 {
			if numList[idx] > targetNum {
				continue
			}

			subResNum := make([]int, 0)

			subNumList := make([]int, len(numList)-idx-1)
			copy(subNumList, numList[idx+1:])

			for counter := 1; counter <= numCounterMap[numList[idx]]; counter += 1 {
				subResNum = append(subResNum, numList[idx])

				if targetNum-counter*numList[idx] == 0 {
					resList = append(resList, append([]int{}, subResNum...))
					break
				} else if targetNum-counter*numList[idx] < 0 {
					break
				}

				subResList := findNum(subNumList, targetNum-counter*numList[idx])
				for _, currentSubRes := range subResList {
					resList = append(resList, append(append([]int{}, subResNum...), currentSubRes...))
				}
			}
		}

		return resList
	}

	return findNum(uniqueCandidateList, target)
}

// @lc code=end

