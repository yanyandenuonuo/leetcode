/*
 * @lc app=leetcode.cn id=1302 lang=golang
 *
 * [1302] 层数最深叶子节点的和
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
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	currentQueue := make([]*TreeNode, 0)
	currentQueue = append(currentQueue, root)
	sum := 0
	for len(currentQueue) > 0 {
		nextQueue := make([]*TreeNode, 0, 2*len(currentQueue))
		sum = 0
		for _, node := range currentQueue {
			sum += node.Val
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		currentQueue = nextQueue
	}

	return sum
}

// @lc code=end

