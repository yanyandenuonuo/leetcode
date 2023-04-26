/*
 * @lc app=leetcode.cn id=1028 lang=golang
 *
 * [1028] 从先序遍历还原二叉树
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
func recoverFromPreorder(traversal string) *TreeNode {
	var dfs func(levelIdx int) *TreeNode
	dfs = func(levelIdx int) *TreeNode {
		// 查找前缀，若不匹配则返回nil
		findCount := 0
		for idx := 0; idx < len(traversal); idx += 1 {
			if traversal[idx] != '-' {
				break
			}
			findCount += 1
		}
		if findCount != levelIdx {
			return nil
		}

		val := 0
		nextIdx := findCount
		for ; nextIdx < len(traversal); nextIdx += 1 {
			if traversal[nextIdx] == '-' {
				break
			}
			val = val*10 + int(traversal[nextIdx]-'0')
		}

		traversal = traversal[nextIdx:]

		return &TreeNode{
			Val:   val,
			Left:  dfs(levelIdx + 1),
			Right: dfs(levelIdx + 1),
		}
	}
	return dfs(0)
}

// @lc code=end

