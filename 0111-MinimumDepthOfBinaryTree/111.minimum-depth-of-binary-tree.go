/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (37.67%)
 * Likes:    1731
 * Dislikes: 757
 * Total Accepted:    451.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 * 
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
 * 
 * Note: A leaf is a node with no children.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [3,9,20,null,null,15,7]
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [2,null,3,null,4,null,5,null,6]
 * Output: 5
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the tree is in the range [0, 10^4].
 * -1000 <= Node.val <= 1000
 * 
 * 
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
func minDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	queueList := make([]*TreeNode, 0)

	queueList = append(queueList, root)
	depth := 1

	for len(queueList) > 0 {
		currentQueue := len(queueList)
		for idx := 0; idx < currentQueue; idx += 1 {
			currentNode := queueList[idx]

			if currentNode.Left == nil && currentNode.Right == nil {
				return depth
			}

			if currentNode.Left != nil {
				queueList = append(queueList, currentNode.Left)
			}
			if currentNode.Right != nil {
				queueList = append(queueList, currentNode.Right)
			}
		}
		if len(queueList) > currentQueue {
			queueList = queueList[currentQueue:]
		} else {
			queueList = make([]*TreeNode, 0)
		}
		depth += 1
	}

	return depth
}
// @lc code=end

