/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
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
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs func(treeNode *TreeNode) int
	dfs = func(treeNode *TreeNode) int {
		if treeNode == nil {
			return 0
		}

		leftSum := dfs(treeNode.Left)
		rightSum := dfs(treeNode.Right)

		diff := leftSum - rightSum
		if diff < 0 {
			diff = -diff
		}
		res += diff

		return treeNode.Val + leftSum + rightSum
	}

	dfs(root)

	return res
}

// @lc code=end

