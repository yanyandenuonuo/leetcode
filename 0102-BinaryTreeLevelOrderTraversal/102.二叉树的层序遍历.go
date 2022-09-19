/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	resList := make([][]int, 0)

	nodeQueue := make([]*TreeNode, 0)
	nodeQueue = append(nodeQueue, root)
	for level := 0; len(nodeQueue) > 0; level += 1 {
		curLevelList := make([]int, 0, len(nodeQueue))
		nextQueue := make([]*TreeNode, 0)

		for _, curNode := range nodeQueue {
			curLevelList = append(curLevelList, curNode.Val)
			if curNode.Left != nil {
				nextQueue = append(nextQueue, curNode.Left)
			}
			if curNode.Right != nil {
				nextQueue = append(nextQueue, curNode.Right)
			}
		}
		resList = append(resList, curLevelList)
		nodeQueue = nextQueue
	}

	return resList
}

// @lc code=end

