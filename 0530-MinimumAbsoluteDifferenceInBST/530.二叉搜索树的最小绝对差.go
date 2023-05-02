/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	// solution1: 中序遍历获得数组，然后数组相邻数字之间最小差
	/**
	nodeList := inOrder(root)
	minDiff := 1<<31 - 1
	for idx := 1; idx < len(nodeList); idx += 1 {
		if nodeList[idx]-nodeList[idx-1] < minDiff {
			minDiff = nodeList[idx] - nodeList[idx-1]
		}
	}

	return minDiff
	*/

	// solution2: DFS，直接求diff值
	minDiff, preVal := 1<<31-1, 1<<31-1
	var dfs func(treeNode *TreeNode)
	dfs = func(treeNode *TreeNode) {
		if treeNode == nil {
			return
		}
		dfs(treeNode.Left)
		if treeNode.Val > preVal && treeNode.Val-preVal < minDiff {
			minDiff = treeNode.Val - preVal
		}
		preVal = treeNode.Val
		dfs(treeNode.Right)
	}

	dfs(root)
	return minDiff
}

func inOrder(treeNode *TreeNode) []int {
	if treeNode == nil {
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, inOrder(treeNode.Left)...)
	res = append(res, treeNode.Val)
	res = append(res, inOrder(treeNode.Right)...)
	return res
}

// @lc code=end

