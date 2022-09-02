/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 */

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
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
			continue
		} else if firstList[idxF][0] >= secondList[idxS][0] && firstList[idxF][1] <= secondList[idxS][1] {
			res = append(res, []int{firstList[idxF][0], firstList[idxF][1]})
			idxF += 1
			continue
		} else if firstList[idxF][0] <= secondList[idxS][0] && firstList[idxF][1] <= secondList[idxS][1] {
			res = append(res, []int{secondList[idxS][0], firstList[idxF][1]})
			idxF += 1
			continue
		} else if firstList[idxF][0] >= secondList[idxS][0] && firstList[idxF][1] >= secondList[idxS][1] {
			res = append(res, []int{firstList[idxF][0], secondList[idxS][1]})
			idxS += 1
			continue
		}
	}

	return res
}

// @lc code=end

