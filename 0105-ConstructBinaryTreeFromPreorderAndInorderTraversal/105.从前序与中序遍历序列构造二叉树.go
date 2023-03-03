/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 前序遍历： [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
	// 中序遍历： [ [左子树的前序遍历结果], 根节点, [右子树的前序遍历结果] ]
	if len(preorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = preorder[0]
	rootIdx := 0
	for idx, val := range inorder {
		if val == root.Val {
			rootIdx = idx
			break
		}
	}
	root.Left = buildTree(preorder[1:rootIdx+1], inorder[:rootIdx])
	root.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
	return root
}

// @lc code=end

