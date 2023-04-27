/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
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
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	targetVal := root.Val
	currentQueue := make([]*TreeNode, 0)
	currentQueue = append(currentQueue, root)

	for len(currentQueue) > 0 {
		nextQueue := make([]*TreeNode, 0, 2*len(currentQueue))
		for _, treeNode := range currentQueue {
			if treeNode.Val != root.Val && (targetVal == root.Val || treeNode.Val < targetVal) {
				targetVal = treeNode.Val
			}

			if treeNode.Left != nil {
				nextQueue = append(nextQueue, treeNode.Left)
			}

			if treeNode.Right != nil {
				nextQueue = append(nextQueue, treeNode.Right)
			}
		}
		currentQueue = nextQueue
	}

	if targetVal != root.Val {
		return targetVal
	}

	return -1
}

// @lc code=end

