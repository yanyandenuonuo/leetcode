/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

// @lc code=start
func kClosest(points [][]int, k int) [][]int {
	// solution1: quicksort
	/**
	distanceList := make([][]int, len(points))
	for idx, pointIdx := range points {
		distanceList[idx] = []int{idx, pointIdx[0]*pointIdx[0] + pointIdx[1]*pointIdx[1]}
	}

	quicksortK(distanceList, 0, len(distanceList)-1, k)
	res := make([][]int, 0, k)
	for idx := 0; idx < k; idx += 1 {
		res = append(res, points[distanceList[idx][0]])
	}
	return res
	*/

	// solution2: heap
	distanceList := make([][]int, k)
	for idx, pointIdx := range points {
		if idx < k {
			distanceList[idx] = []int{idx, pointIdx[0]*pointIdx[0] + pointIdx[1]*pointIdx[1]}
			continue
		} else if idx == k {
			heapSort(distanceList)
		}
		distance := pointIdx[0]*pointIdx[0] + pointIdx[1]*pointIdx[1]
		if distance < distanceList[0][1] {
			distanceList[0] = []int{idx, distance}
			heapSort(distanceList)
		}
	}
	res := make([][]int, 0, k)
	for idx := 0; idx < k; idx += 1 {
		res = append(res, points[distanceList[idx][0]])
	}
	return res
}

func heapSort(distanceList [][]int) {
	for idx := len(distanceList)/2 - 1; idx >= 0; idx -= 1 {
		heapify(distanceList, idx, len(distanceList))
	}
	for idx := len(distanceList) - 1; idx > 0; idx -= 1 {
		distanceList[idx], distanceList[0] = distanceList[0], distanceList[idx]
		heapify(distanceList, 0, idx)
	}
}

func heapify(distanceList [][]int, idx, length int) {
	leftChildIdx := 2*idx + 1
	rightChildIdx := 2*idx + 2

	minimalIdx := idx

	if leftChildIdx < length && distanceList[leftChildIdx][1] < distanceList[minimalIdx][1] {
		minimalIdx = leftChildIdx
	}

	if rightChildIdx < length && distanceList[rightChildIdx][1] < distanceList[rightChildIdx][1] {
		minimalIdx = rightChildIdx
	}
	if minimalIdx != idx {
		distanceList[minimalIdx], distanceList[idx] = distanceList[idx], distanceList[minimalIdx]
		heapify(distanceList, minimalIdx, length)
	}
}

func quicksortK(distanceList [][]int, leftIdx, rightIdx, k int) {
	if leftIdx < rightIdx {
		pivotIdx := findPivotIdx(distanceList, leftIdx, rightIdx)
		if pivotIdx == k-1 {
			return
		} else if pivotIdx < k-1 {
			quicksortK(distanceList, pivotIdx+1, rightIdx, k)
		} else if pivotIdx > k-1 {
			quicksortK(distanceList, leftIdx, pivotIdx-1, k)
		}
	}

	return
}

func findPivotIdx(distanceList [][]int, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx+1)
	distanceList[pivotIdx], distanceList[rightIdx] = distanceList[rightIdx], distanceList[pivotIdx]

	swapIdx := leftIdx

	for idx := leftIdx; idx < rightIdx; idx += 1 {
		if distanceList[idx][1] <= distanceList[rightIdx][1] {
			distanceList[swapIdx], distanceList[idx] = distanceList[idx], distanceList[swapIdx]
			swapIdx += 1
		}
	}
	distanceList[swapIdx], distanceList[rightIdx] = distanceList[rightIdx], distanceList[swapIdx]
	return swapIdx
}

// @lc code=end

