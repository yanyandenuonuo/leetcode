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
	// solution2: 从根节点出发，直接判断当前节点是否为p, q的祖先，若是则找下一个，最后一个符合的即为最近的公共祖先
	// if p.Val > q.Val {
	// 	p, q = q, p
	// }

	// for {
	// 	if p.Val < root.Val && q.Val > root.Val {
	// 		return root
	// 	} else if p.Val < root.Val && q.Val < root.Val {
	// 		root = root.Left
	// 	} else if p.Val > root.Val && q.Val > root.Val {
	// 		root = root.Right
	// 	} else if p.Val == root.Val || q.Val == root.Val {
	// 		return root
	// 	}
	// }
	// return nil

	// solution3: 从根节点出发，直接判断p, q节点在当前节点左右两侧，q, q其中一个节点为当前节点，还是同一侧
	//				在当前节点左右两侧，则(root.Val-p.Val)*(root.Val-q.Val) < 0，那么当前节点为最近公共祖先
	//				当前节点为其中一个节点，则(root.Val-p.Val)*(root.Val-q.Val) == 0，那么当前节点为最近公共祖先
	//				在当前节点同一侧，则(root.Val-p.Val)*(root.Val-q.Val) > 0，继续进行搜索
	for (root.Val-p.Val)*(root.Val-q.Val) > 0 {
		if root.Val > p.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// @lc code=end

