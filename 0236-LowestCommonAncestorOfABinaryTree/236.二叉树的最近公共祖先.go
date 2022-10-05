/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// solution1: hashmap, 找出所有节点父节点，再遍历p节点到root节点并做记忆hashmap，再从q节点遍历到root节点，若浏览到的节点被记忆，则该节点即为最近公共祖先
	// solution2: 分别找出root到p, q节点的路径，再2个路径中找分叉点即可
	// solution3: 递归搜索：较为抽象，可以分2种情况考虑，
	// 			  1. 搜索节点为p或q，那么该节点必为最近的公共祖先（题目有给前提说p，q均存在于树中），
	//			  2. 节点的left，right均搜索到值，那该节点必为最近的公共祖先，否则最近的公共祖在左子节点或右子结点
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// @lc code=end

