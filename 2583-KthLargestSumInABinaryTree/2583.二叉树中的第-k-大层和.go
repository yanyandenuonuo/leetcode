/*
 * @lc app=leetcode.cn id=2583 lang=golang
 *
 * [2583] 二叉树中的第 K 大层和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}
	currentQueue := make([]*TreeNode, 0)
	currentQueue = append(currentQueue, root)
	res := make([]int64, 0)
	for idx := 1; len(currentQueue) > 0; idx += 1 {
		nextQueue := make([]*TreeNode, 0, 2*len(currentQueue))
		sum := int64(0)
		for _, node := range currentQueue {
			sum += int64(node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		res = append(res, sum)
		currentQueue = nextQueue
	}

	if len(res) < k {
		return -1
	}

	return findKthLargest(res, k)
}

func findKthLargest(nums []int64, k int) int64 {
	k = len(nums) - k
	quickSort(nums, 0, len(nums)-1, k)
	return nums[k]
}

func quickSort(nums []int64, leftIdx, rightIdx, k int) {
	if leftIdx >= rightIdx {
		return
	}
	pivotIdx := partition(nums, leftIdx, rightIdx)
	if pivotIdx > k {
		quickSort(nums, leftIdx, pivotIdx-1, k)
	} else if pivotIdx < k {
		quickSort(nums, pivotIdx+1, rightIdx, k)
	} else {
		return
	}
}

func partition(nums []int64, leftIdx, rightIdx int) int {
	pivotIdx := leftIdx + rand.Intn(rightIdx-leftIdx)
	// swap pivot with right
	nums[rightIdx], nums[pivotIdx] = nums[pivotIdx], nums[rightIdx]
	swapIdx := leftIdx

	for idx := leftIdx; idx < rightIdx; idx += 1 {
		if nums[idx] < nums[rightIdx] {
			nums[idx], nums[swapIdx] = nums[swapIdx], nums[idx]
			swapIdx += 1
		}
	}
	nums[swapIdx], nums[rightIdx] = nums[rightIdx], nums[swapIdx]

	return swapIdx
}

// @lc code=end

