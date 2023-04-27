/*
 * @lc app=leetcode.cn id=2331 lang=golang
 *
 * [2331] 计算布尔二叉树的值
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
func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var dfs func(treeNode *TreeNode) bool
	dfs = func(treeNode *TreeNode) bool {
		if treeNode == nil {
			return false
		}

		// 是叶子节点
		if treeNode.Left == nil && treeNode.Right == nil {
			// 0表示false，1表示true
			return treeNode.Val == 1
		}

		leftVal := dfs(treeNode.Left)
		rightVal := dfs(treeNode.Right)

		// 2表示or，3表示and
		if treeNode.Val == 2 {
			return leftVal || rightVal
		}
		return leftVal && rightVal
	}

	return dfs(root)
}

// @lc code=end

