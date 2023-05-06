/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 */

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	// solution1: 双指针+区间对比
	/**
	res := make([][]int, 0, len(firstList))
	for idxF, idxS := 0, 0; idxF < len(firstList) && idxS < len(secondList); {
		if firstList[idxF][1] < secondList[idxS][0] {
			idxF += 1
			continue
		}

		if firstList[idxF][0] > secondList[idxS][1] {
			idxS += 1
			continue
		}

		if firstList[idxF][0] <= secondList[idxS][0] && firstList[idxF][1] >= secondList[idxS][1] {
			res = append(res, []int{secondList[idxS][0], secondList[idxS][1]})
			idxS += 1
		} else if firstList[idxF][0] >= secondList[idxS][0] && firstList[idxF][1] <= secondList[idxS][1] {
			res = append(res, []int{firstList[idxF][0], firstList[idxF][1]})
			idxF += 1
		} else if firstList[idxF][0] <= secondList[idxS][0] && firstList[idxF][1] <= secondList[idxS][1] {
			res = append(res, []int{secondList[idxS][0], firstList[idxF][1]})
			idxF += 1
		} else if firstList[idxF][0] >= secondList[idxS][0] && firstList[idxF][1] >= secondList[idxS][1] {
			res = append(res, []int{firstList[idxF][0], secondList[idxS][1]})
			idxS += 1
		}
	}

	return res
	*/

	// solution2: 双指针+归并区间
	res := make([][]int, 0, len(firstList))
	for idxF, idxS := 0, 0; idxF < len(firstList) && idxS < len(secondList); {
		leftIdx := firstList[idxF][0]
		if secondList[idxS][0] > leftIdx {
			leftIdx = secondList[idxS][0]
		}

		rightIdx := firstList[idxF][1]
		if secondList[idxS][1] < rightIdx {
			rightIdx = secondList[idxS][1]
		}

		if rightIdx >= leftIdx {
			res = append(res, []int{leftIdx, rightIdx})
		}

		if firstList[idxF][1] < secondList[idxS][1] {
			idxF += 1
		} else {
			idxS += 1
		}
	}

	return res
}

// @lc code=end

