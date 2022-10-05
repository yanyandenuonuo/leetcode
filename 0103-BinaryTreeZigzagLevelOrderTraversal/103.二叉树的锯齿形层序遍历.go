/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)

	currentQueue := make([]*TreeNode, 0)
	currentQueue = append(currentQueue, root)
	currentDirect := 1

	for len(currentQueue) > 0 {
		nodeCount := len(currentQueue)
		idx := nodeCount - 1

		nextQueue := make([]*TreeNode, 0, 2*nodeCount)
		currentRes := make([]int, 0, nodeCount)
		for idx >= 0 {
			currentRes = append(currentRes, currentQueue[idx].Val)
			if currentDirect == 1 {
				if currentQueue[idx].Left != nil {
					nextQueue = append(nextQueue, currentQueue[idx].Left)
				}
				if currentQueue[idx].Right != nil {
					nextQueue = append(nextQueue, currentQueue[idx].Right)
				}
			} else if currentDirect == -1 {
				if currentQueue[idx].Right != nil {
					nextQueue = append(nextQueue, currentQueue[idx].Right)
				}
				if currentQueue[idx].Left != nil {
					nextQueue = append(nextQueue, currentQueue[idx].Left)
				}
			}
			idx -= 1
		}
		currentDirect = -currentDirect
		res = append(res, currentRes)
		currentQueue = nextQueue
	}
	return res
}

// @lc code=end

