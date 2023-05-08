/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	// solution1: DFS
	var dfs func(treeNode *TreeNode, preSum int) int
	dfs = func(treeNode *TreeNode, preSum int) int {
		if treeNode == nil {
			return 0
		}

		preSum = preSum*10 + treeNode.Val

		if treeNode.Left == nil && treeNode.Right == nil {
			return preSum
		}

		return dfs(treeNode.Left, preSum) + dfs(treeNode.Right, preSum)
	}

	return dfs(root, 0)

	// solution2: BFS
}

// @lc code=end

