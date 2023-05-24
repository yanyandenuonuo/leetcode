/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
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
func minDiffInBST(root *TreeNode) int {
	// solution1: 中序遍历获得数组，然后数组相邻数字之间最小差
	// solution2: DFS，直接求diff值
	minDiff, preVal := 1<<31-1, -1<<31
	var dfs func(treeNode *TreeNode)
	dfs = func(treeNode *TreeNode) {
		if treeNode == nil {
			return
		}

		dfs(treeNode.Left)

		if treeNode.Val > preVal && treeNode.Val-preVal < minDiff {
			minDiff = treeNode.Val - preVal
		}

		preVal = treeNode.Val

		dfs(treeNode.Right)
	}

	dfs(root)

	return minDiff
}

// @lc code=end

