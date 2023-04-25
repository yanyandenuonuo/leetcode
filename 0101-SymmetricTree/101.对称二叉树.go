/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	// solution1: 逐层扫描是否对层
	// solution2: 递归判断两个节点是否对称
	//			  两个节点值相同 && 左节点的左子节点和右节点的右子节点对称 && 左节点的右子节点和右节点的左子节点对称
	if root == nil {
		return true
	}
	return isSymmetricWithTwoNode(root.Left, root.Right)
}

func isSymmetricWithTwoNode(leftNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode == nil || rightNode == nil {
		return false
	}
	return leftNode.Val == rightNode.Val && isSymmetricWithTwoNode(leftNode.Left, rightNode.Right) &&
		isSymmetricWithTwoNode(leftNode.Right, rightNode.Left)
}

// @lc code=end

