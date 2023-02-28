/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	valFrequency := make(map[int]int, len(nums))

	// 统计频率
	for _, val := range nums {
		valFrequency[val] += 1
	}

	frequencyList := make([][]int, 0, k)
	for val, frequency := range valFrequency {
		if len(frequencyList) < k {
			frequencyList = append(frequencyList, []int{frequency, val})
			continue
		} else if len(frequencyList) == k {
			// 构造小顶堆
			for idx := k/2 - 1; idx >= 0; idx -= 1 {
				heapify(frequencyList, idx, k)
			}
		}
		if frequency < frequencyList[0][0] {
			continue
		}
		frequencyList[0] = []int{frequency, val}
		heapify(frequencyList, 0, k)
	}

	res := make([]int, 0, k)
	for _, frequencyInfo := range frequencyList {
		res = append(res, frequencyInfo[1])
	}
	return res
}

func heapify(nums [][]int, idx, length int) {
	leftChild := 2*idx + 1
	rightChild := 2*idx + 2
	lowest := idx
	if leftChild < length && nums[leftChild][0] < nums[lowest][0] {
		lowest = leftChild
	}
	if rightChild < length && nums[rightChild][0] < nums[lowest][0] {
		lowest = rightChild
	}
	if lowest != idx {
		nums[idx], nums[lowest] = nums[lowest], nums[idx]
		// 交换完成之后需要对lowest节点重新进行heapify
		heapify(nums, lowest, length)
	}
}

// @lc code=end

