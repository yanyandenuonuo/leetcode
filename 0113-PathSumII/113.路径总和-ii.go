/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	resList := make([][]int, 0)
	pathList := make([]int, 0)

	var hasPathSum func(rootNode *TreeNode, targetNum int)
	hasPathSum = func(rootNode *TreeNode, targetNum int) {
		if rootNode == nil {
			return
		}

		if rootNode.Left == nil && rootNode.Right == nil {
			if rootNode.Val == targetNum {
				pathList = append(pathList, rootNode.Val)
				resList = append(resList, append([]int{}, pathList...))
				pathList = pathList[:len(pathList)-1]
			}
			return
		}

		pathList = append(pathList, rootNode.Val)

		// 遍历左子节点
		hasPathSum(rootNode.Left, targetNum-rootNode.Val)

		// 遍历右子节点
		hasPathSum(rootNode.Right, targetNum-rootNode.Val)
		pathList = pathList[:len(pathList)-1]
	}
	hasPathSum(root, targetSum)
	return resList
}

// @lc code=end

