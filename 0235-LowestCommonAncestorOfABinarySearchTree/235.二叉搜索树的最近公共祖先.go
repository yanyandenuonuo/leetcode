/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// solution1: 分别找出从根节点出发到达p q的路径，然后比较路径，分叉点之前的最后一个即为最近的公共祖先
	// solution2: 从根节点出发，直接判断当前节点是否为p q的祖先，若是则找下一个，最后一个符合的即为最近的公共祖先
	if p.Val > q.Val {
		p, q = q, p
	}

	for {
		if p.Val < root.Val && q.Val > root.Val {
			return root
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else if p.Val == root.Val || q.Val == root.Val {
			return root
		}
	}
	return nil
}

// @lc code=end

