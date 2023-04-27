/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := make([]string, 0)
	pathList := make([]string, 0)

	var dfs func(treeNode *TreeNode)
	dfs = func(treeNode *TreeNode) {
		if treeNode.Left == nil && treeNode.Right == nil {
			res = append(res, strings.Join(append(pathList, strconv.Itoa(treeNode.Val)), "->"))
			return
		}
		pathList = append(pathList, strconv.Itoa(treeNode.Val))
		if treeNode.Left != nil {
			dfs(treeNode.Left)
		}
		if treeNode.Right != nil {
			dfs(treeNode.Right)
		}
		pathList = pathList[:len(pathList)-1]
	}

	dfs(root)

	return res
}

// @lc code=end

