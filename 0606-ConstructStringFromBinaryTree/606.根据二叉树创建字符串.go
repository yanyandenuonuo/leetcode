/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
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
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left != nil && root.Right != nil {
		return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")" + "(" + tree2str(root.Right) + ")"
	} else if root.Left != nil && root.Right == nil {
		return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")"
	} else if root.Left == nil && root.Right != nil {
		return strconv.Itoa(root.Val) + "()" + "(" + tree2str(root.Right) + ")"
	} else {
		return strconv.Itoa(root.Val)
	}
}

// @lc code=end

