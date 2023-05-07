/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	// solution1: DFS
	if root == nil {
		return 0
	}
	maxHeight := treeHeight(root)
	return maxHeight

	// solution2: BFS
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := treeHeight(root.Left)
	rightHeight := treeHeight(root.Right)

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// @lc code=end

