/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] 叶子相似的树
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// solution:dfs

	root1Leaf := getTreeLeaf(root1)
	root2Leaf := getTreeLeaf(root2)

	if len(root1Leaf) != len(root2Leaf) {
		return false
	}

	for idx, val := range root1Leaf {
		if val != root2Leaf[idx] {
			return false
		}
	}

	return true
}

func getTreeLeaf(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	treeLeaf := make([]int, 0)

	if root.Left == nil && root.Right == nil {
		treeLeaf = append(treeLeaf, root.Val)
	} else {
		if root.Left != nil {
			treeLeaf = append(treeLeaf, getTreeLeaf(root.Left)...)
		}

		if root.Right != nil {
			treeLeaf = append(treeLeaf, getTreeLeaf(root.Right)...)
		}
	}

	return treeLeaf
}

// @lc code=end

