/*
 * @lc app=leetcode.cn id=958 lang=golang
 *
 * [958] 二叉树的完全性检验
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
func isCompleteTree(root *TreeNode) bool {
	// solution1: BFS，给节点编号，看节点值是否连续
	// solution1: BFS，层序遍历判断空值位置
	if root == nil {
		return true
	}

	nodeQueue := make([]*TreeNode, 0)
	nodeQueue = append(nodeQueue, root)

	findNil := false
	for len(nodeQueue) > 0 {
		nextNodeQueue := make([]*TreeNode, 0, 2*len(nodeQueue))
		for _, node := range nodeQueue {
			if node.Left == nil {
				findNil = true
			} else {
				if findNil {
					return false
				}

				nextNodeQueue = append(nextNodeQueue, node.Left)
			}

			if node.Right == nil {
				findNil = true
			} else {
				if findNil {
					return false
				}

				nextNodeQueue = append(nextNodeQueue, node.Right)
			}
		}

		nodeQueue = nextNodeQueue
	}

	return true
}

// @lc code=end

