/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] 另一棵树的子树
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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// solution1: 深度优先搜索+暴力匹配
	if root == nil {
		return false
	}

	return isSubTreeOfCurrentNode(root, subRoot) || isSubtree(root.Left, subRoot) ||
		isSubtree(root.Right, subRoot)

	// solution2: 深度优先搜索+(KMP或者Rabin-Karp)
	// todo
	// solution3: 树hash
	// todo
}

func isSubTreeOfCurrentNode(treeNode, subNode *TreeNode) bool {
	if treeNode == nil && subNode == nil {
		return true
	}

	if treeNode == nil || subNode == nil {
		return false
	}

	return treeNode.Val == subNode.Val && isSubTreeOfCurrentNode(treeNode.Left, subNode.Left) &&
		isSubTreeOfCurrentNode(treeNode.Right, subNode.Right)
}

// @lc code=end

