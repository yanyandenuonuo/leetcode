/*
 * @lc app=leetcode.cn id=997 lang=golang
 *
 * [997] 找到小镇的法官
 */

// @lc code=start
func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}

	trustMap, trustedMap, candidateList := make(map[int]map[int]bool, n), make(map[int]map[int]bool, n), make([]int, 0)
	for _, item := range trust {
		if _, isExist := trustMap[item[0]]; !isExist {
			trustMap[item[0]] = make(map[int]bool, n)
		}

		trustMap[item[0]][item[1]] = true

		if _, isExist := trustedMap[item[1]]; !isExist {
			trustedMap[item[1]] = make(map[int]bool, n)
		}

		trustedMap[item[1]][item[0]] = true

		if len(trustedMap[item[1]]) == n-1 {
			candidateList = append(candidateList, item[1])
		}
	}

	judgeIdx := -1

	for _, candidateIdx := range candidateList {
		if len(trustMap[candidateIdx]) > 0 {
			continue
		}

		if judgeIdx != -1 {
			judgeIdx = -1
			break
		}

		judgeIdx = candidateIdx
	}

	return judgeIdx
}

// @lc code=end

