/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	// solution: DFS
	maxSum := -1 << 31

	var getNodeMaxPath func(treeNode *TreeNode) int
	getNodeMaxPath = func(treeNode *TreeNode) int {
		if treeNode == nil {
			return 0
		}

		leftNodePath := getNodeMaxPath(treeNode.Left)
		if leftNodePath < 0 {
			leftNodePath = 0
		}

		rightNodePath := getNodeMaxPath(treeNode.Right)
		if rightNodePath < 0 {
			rightNodePath = 0
		}

		nodePath := treeNode.Val + leftNodePath + rightNodePath

		if nodePath > maxSum {
			maxSum = nodePath
		}

		if leftNodePath > rightNodePath {
			return treeNode.Val + leftNodePath
		} else {
			return treeNode.Val + rightNodePath
		}
	}

	getNodeMaxPath(root)

	return maxSum
}

// @lc code=end

