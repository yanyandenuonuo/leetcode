/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (48.11%)
 * Likes:    2110
 * Dislikes: 39
 * Total Accepted:    258.1K
 * Total Submissions: 536.3K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * Given inorder and postorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * inorder = [9,3,15,20,7]
 * postorder = [9,15,7,20,3]
 *
 * Return the following binary tree:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 中序遍历：[左子节点] [节点] [右子节点]
	// 后序遍历：[左子节点] [右子节点] [节点]
	if len(inorder) == 0 {
		return nil
	} else if len(inorder) == 1 {
		return &TreeNode{
			Val:   inorder[0],
			Left:  nil,
			Right: nil,
		}
	}

	idx := 0
	for ; idx < len(inorder); idx += 1 {
		if inorder[idx] == postorder[len(postorder)-1] {
			break
		}
	}
	return &TreeNode{
		Val:   inorder[idx],
		Left:  buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
	}
}

// @lc code=end

