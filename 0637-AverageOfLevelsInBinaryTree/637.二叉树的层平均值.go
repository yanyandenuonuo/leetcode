/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	currentQueue := make([]*TreeNode, 0)
	currentQueue = append(currentQueue, root)
	res := make([]float64, 0)
	for len(currentQueue) > 0 {
		nextQueue := make([]*TreeNode, 0, 2*len(currentQueue))
		sum := 0
		for _, node := range currentQueue {
			sum += node.Val
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(currentQueue)))
		currentQueue = nextQueue
	}

	return res
}

// @lc code=end

