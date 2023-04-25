/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] 根据前序和后序遍历构造二叉树
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	// 前序遍历：[节点] [左子节点] [右子节点]
	// 后序遍历：[左子节点] [右子节点] [节点]
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
	}

	idx := 0
	for ; idx < len(postorder); idx += 1 {
		if postorder[idx] == preorder[1] {
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  constructFromPrePost(preorder[1:idx+2], postorder[:idx+1]),
		Right: constructFromPrePost(preorder[idx+2:], postorder[idx+1:len(postorder)-1]),
	}
}

// @lc code=end

