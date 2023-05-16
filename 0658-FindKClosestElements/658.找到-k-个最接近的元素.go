/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	// solution1: 快排
	// solution2: 堆排序
	// solution3: 二分查找+双指针
	leftIdx, rightIdx := 0, len(arr)-1
	for leftIdx < rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if arr[midIdx] >= x {
			rightIdx = midIdx
		} else {
			leftIdx = midIdx + 1
		}
	}

	// 注意这里取值的边界问题
	rightIdx = leftIdx
	leftIdx -= 1

	// 查找leftIdx和rightIdx可以用库函数
	// rightIdx := sort.SearchInts(arr, x)
	// leftIdx := rightIdx - 1

	for counter := 0; counter < k; counter += 1 {
		if leftIdx < 0 {
			rightIdx += 1
		} else if rightIdx >= len(arr) {
			leftIdx -= 1
		} else {
			if arr[rightIdx]-x < x-arr[leftIdx] {
				rightIdx += 1
			} else {
				leftIdx -= 1
			}
		}
	}

	return arr[leftIdx+1 : rightIdx]
}

// @lc code=end

