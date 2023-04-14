/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	return treeHeight(root) >= 0
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := treeHeight(root.Left)
	rightHeight := treeHeight(root.Right)

	if leftHeight == -1 || rightHeight == -1 || math.Abs(float64(rightHeight-leftHeight)) > 1 {
		return -1
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// @lc code=end

