/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, maxLength := findMaxLength(root, 0)
	return maxLength
}

func findMaxLength(root *TreeNode, maxLength int) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftHeight, leftLength := findMaxLength(root.Left, maxLength)
	rightHeight, rightLength := findMaxLength(root.Right, maxLength)

	if leftLength > maxLength {
		maxLength = leftLength
	}

	if rightLength > maxLength {
		maxLength = rightLength
	}

	if leftHeight+rightHeight > maxLength {
		maxLength = leftHeight + rightHeight
	}

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1, maxLength
}

// @lc code=end

