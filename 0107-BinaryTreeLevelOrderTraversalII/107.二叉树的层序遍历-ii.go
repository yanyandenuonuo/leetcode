/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	currentQueue := make([]*TreeNode, 0)
	currentQueue = append(currentQueue, root)
	resList := make([][]int, 0)
	for len(currentQueue) > 0 {
		nextQueue := make([]*TreeNode, 0, 2*len(currentQueue))
		res := make([]int, 0)

		for _, node := range currentQueue {
			res = append(res, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		resList = append([][]int{res}, resList...)
		currentQueue = nextQueue
	}

	return resList
}

// @lc code=end

