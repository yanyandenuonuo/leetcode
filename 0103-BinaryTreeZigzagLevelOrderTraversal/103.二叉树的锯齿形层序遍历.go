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
	currentDirect := true

	for len(currentQueue) > 0 {
		currentRes := make([]int, 0, len(currentQueue))
		nextQueue := make([]*TreeNode, 0, 2*len(currentQueue))

		for idx := len(currentQueue) - 1; idx >= 0; idx -= 1 {
			currentRes = append(currentRes, currentQueue[idx].Val)
			if currentDirect {
				if currentQueue[idx].Left != nil {
					nextQueue = append(nextQueue, currentQueue[idx].Left)
				}
				if currentQueue[idx].Right != nil {
					nextQueue = append(nextQueue, currentQueue[idx].Right)
				}
			} else {
				if currentQueue[idx].Right != nil {
					nextQueue = append(nextQueue, currentQueue[idx].Right)
				}
				if currentQueue[idx].Left != nil {
					nextQueue = append(nextQueue, currentQueue[idx].Left)
				}
			}
		}
		res = append(res, currentRes)

		currentQueue = nextQueue
		currentDirect = !currentDirect
	}
	return res
}

// @lc code=end

