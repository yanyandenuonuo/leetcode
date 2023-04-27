/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
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
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	currentQueue := make([]*TreeNode, 0)
	currentQueue = append(currentQueue, root)
	for findX, findY := false, false; len(currentQueue) > 0; {
		nextQueue := make([]*TreeNode, 0, 2*len(currentQueue))
		for _, treeNode := range currentQueue {
			if treeNode.Left != nil && treeNode.Right != nil &&
				(treeNode.Left.Val == x && treeNode.Right.Val == y ||
					treeNode.Left.Val == y && treeNode.Right.Val == x) {
				return false
			}

			if treeNode.Left != nil {
				if treeNode.Left.Val == x {
					findX = true
				} else if treeNode.Left.Val == y {
					findY = true
				}
				nextQueue = append(nextQueue, treeNode.Left)
			}

			if treeNode.Right != nil {
				if treeNode.Right.Val == x {
					findX = true
				} else if treeNode.Right.Val == y {
					findY = true
				}
				nextQueue = append(nextQueue, treeNode.Right)
			}

			if findX && findY {
				return true
			}
		}
		if findX || findY {
			return false
		}
		currentQueue = nextQueue
	}
	return false
}

// @lc code=end

