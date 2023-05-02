/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	nodeQueue := make([]*Node, 0)
	nodeQueue = append(nodeQueue, root)
	counter := 0
	for len(nodeQueue) > 0 {
		nextQueue := make([]*Node, 0)
		for _, node := range nodeQueue {
			if len(node.Children) > 0 {
				nextQueue = append(nextQueue, node.Children...)
			}
		}
		counter += 1
		nodeQueue = nextQueue
	}
	return counter
}

// @lc code=end

