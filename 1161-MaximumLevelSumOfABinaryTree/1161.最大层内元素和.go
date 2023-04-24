/*
 * @lc app=leetcode.cn id=1161 lang=golang
 *
 * [1161] 最大层内元素和
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
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	currentQueue := make([]*TreeNode, 0)
	currentQueue = append(currentQueue, root)
	maxSum := -1 << 31
	maxLevel := 0
	for idx := 1; len(currentQueue) > 0; idx += 1 {
		nextQueue := make([]*TreeNode, 0, 2*len(currentQueue))
		sum := 0
		for _, node := range currentQueue {
			sum += node.Val
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}

		if sum > maxSum {
			maxSum = sum
			maxLevel = idx
		}

		currentQueue = nextQueue

	}

	return maxLevel
}

// @lc code=end

