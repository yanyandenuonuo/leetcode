/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	currentQueue := make([]*TreeNode, 0)
	currentQueue = append(currentQueue, root)
	res := make([]int, 0)
	for len(currentQueue) > 0 {
		res = append(res, currentQueue[len(currentQueue)-1].Val)

		nextQueue := make([]*TreeNode, 0, 2*len(currentQueue))

		for _, node := range currentQueue {
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		currentQueue = nextQueue
	}

	return res
}

// @lc code=end

