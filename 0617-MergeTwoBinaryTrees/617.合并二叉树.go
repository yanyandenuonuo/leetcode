/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	root := new(TreeNode)

	if root1 != nil && root2 == nil {
		root.Val = root1.Val

		// left node
		root.Left = mergeTrees(root1.Left, nil)

		// right node
		root.Right = mergeTrees(root1.Right, nil)
	} else if root1 == nil && root2 != nil {
		root.Val = root2.Val

		// left node
		root.Left = mergeTrees(nil, root2.Left)

		// right node
		root.Right = mergeTrees(nil, root2.Right)
	} else {
		root.Val = root1.Val + root2.Val

		// left node
		root.Left = mergeTrees(root1.Left, root2.Left)

		// right node
		root.Right = mergeTrees(root1.Right, root2.Right)
	}

	return root
}

// @lc code=end

