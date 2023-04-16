/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	// mail => accountIdx
	mailIdxMap := make(map[string]int, len(accounts))

	// mergedAccountIdx => realAccountIdx
	mergeToIdxMap := make(map[int]int, len(accounts))

	// realAccountIdx => email
	accountMap := make(map[int]map[string]bool, len(accounts))

	for accountIdx, account := range accounts {
		// check if the account need merge, if exist then set the accountIdx with merged Idx
		realMergeToAccountIdx := -1
		for idx := 1; idx < len(account); idx += 1 {
			if sameAccountIdx, isExist := mailIdxMap[account[idx]]; isExist && sameAccountIdx != accountIdx && (realMergeToAccountIdx == -1 || sameAccountIdx != realMergeToAccountIdx) {
				// mail exist, need merge account
				for mergeToAccountIdx, hadBeenMerged := mergeToIdxMap[sameAccountIdx]; hadBeenMerged; mergeToAccountIdx, hadBeenMerged = mergeToIdxMap[sameAccountIdx] {
					// same account had been merged, should find the real merged account
					sameAccountIdx = mergeToAccountIdx
				}

				if realMergeToAccountIdx == -1 {
					mergeToIdxMap[accountIdx] = sameAccountIdx
					realMergeToAccountIdx = sameAccountIdx
				} else if sameAccountIdx != realMergeToAccountIdx {
					// A1 a b
					// A2 c d
					// A3 b c
					// merge sameAccountIdx to realMergeToAccountIdx
					for email := range accountMap[sameAccountIdx] {
						accountMap[realMergeToAccountIdx][email] = true
					}
					mergeToIdxMap[sameAccountIdx] = realMergeToAccountIdx
					delete(accountMap, sameAccountIdx)
				}
			}
			mailIdxMap[account[idx]] = accountIdx
		}

		if realMergeToAccountIdx == -1 {
			realMergeToAccountIdx = accountIdx
		}

		// set email to accountMap
		if _, isExist := accountMap[realMergeToAccountIdx]; !isExist {
			accountMap[realMergeToAccountIdx] = make(map[string]bool)
		}

		for idx := 1; idx < len(account); idx += 1 {
			accountMap[realMergeToAccountIdx][account[idx]] = true
		}
	}

	accountList := make([][]string, 0, len(accountMap))
	for accountIdx, emailMap := range accountMap {
		mailList := make([]string, 0, len(emailMap))
		for email := range emailMap {
			mailList = append(mailList, email)
		}
		// sort mailList
		quickSort(mailList, 0, len(mailList)-1)

		accountList = append(accountList, append([]string{accounts[accountIdx][0]}, mailList...))
	}
	return accountList
}

func quickSort(stringList []string, leftIdx, rightIdx int) {
	if leftIdx >= rightIdx {
		return
	}

	pivotIdx := partition(stringList, leftIdx, rightIdx)
	quickSort(stringList, leftIdx, pivotIdx-1)
	quickSort(stringList, pivotIdx+1, rightIdx)
}

func partition(stringList []string, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx+1)
	// swap pivot with right
	stringList[rightIdx], stringList[pivotIdx] = stringList[pivotIdx], stringList[rightIdx]

	swapIdx := leftIdx
	for idx := leftIdx; idx < rightIdx; idx += 1 {
		if stringList[idx] < stringList[rightIdx] {
			stringList[idx], stringList[swapIdx] = stringList[swapIdx], stringList[idx]
			swapIdx += 1
		}
	}
	stringList[swapIdx], stringList[rightIdx] = stringList[rightIdx], stringList[swapIdx]

	return swapIdx
}

// @lc code=end

