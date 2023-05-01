/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
func findMode(root *TreeNode) []int {
	// solution1: 使用hash表记录频率
	// solution2: 使用中序遍历
	maxFrequence, targetVal, counter := 0, 0, 0
	res := make([]int, 0)
	updateRes := func(val int) {
		if val == targetVal {
			counter += 1
		} else {
			counter = 1
			targetVal = val
		}
		if counter > maxFrequence {
			maxFrequence = counter
			res = []int{val}
		} else if counter == maxFrequence {
			res = append(res, val)
		}
	}

	var dfs func(treeNode *TreeNode)
	dfs = func(treeNode *TreeNode) {
		if treeNode == nil {
			return
		}
		dfs(treeNode.Left)
		updateRes(treeNode.Val)
		dfs(treeNode.Right)
	}

	dfs(root)
	return res

	// solution3: Morris中序遍历
	// todo
}

// @lc code=end

