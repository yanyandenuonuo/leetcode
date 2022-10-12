/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	currentQueue := make([]*Node, 0)
	currentQueue = append(currentQueue, root)
	for len(currentQueue) > 0 {
		nextQuene := make([]*Node, 0)
		currentRes := make([]int, 0, len(currentQueue))
		for _, currentNode := range currentQueue {
			currentRes = append(currentRes, currentNode.Val)
			nextQuene = append(nextQuene, currentNode.Children...)
		}
		res = append(res, currentRes)
		currentQueue = nextQuene
	}

	return res
}

// @lc code=end

