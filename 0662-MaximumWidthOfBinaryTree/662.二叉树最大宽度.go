/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
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
type NodeIdx struct {
	treeNode *TreeNode
	Idx      int
}

func widthOfBinaryTree(root *TreeNode) int {
	// solution1: 利用二叉树性质给节点编号，通过DFS找到每层最左边节点的编号以及最右边节点的编号，宽度即为最右边编号-最左边编号+1
	// solution2: 利用二叉树性质给节点编号，通过BFS找到每层最左边节点的编号以及最右边节点的编号，宽度即为最右边编号-最左边编号+1
	//			  因为中间的空节点也计入宽度，所以没法通过直接计算每层节点数的方式来计算宽度
	if root == nil {
		return 0
	}
	currentQueue := make([]*NodeIdx, 0)
	currentQueue = append(currentQueue, &NodeIdx{root, 0})
	maxWidth := 1

	for len(currentQueue) > 0 {
		currentWidth := currentQueue[len(currentQueue)-1].Idx - currentQueue[0].Idx + 1
		if currentWidth > maxWidth {
			maxWidth = currentWidth
		}

		nextQueue := make([]*NodeIdx, 0, 2*len(currentQueue))
		for _, nodeIdx := range currentQueue {
			if nodeIdx.treeNode.Left != nil {
				nextQueue = append(nextQueue, &NodeIdx{nodeIdx.treeNode.Left, 2*nodeIdx.Idx + 1})
			}

			if nodeIdx.treeNode.Right != nil {
				nextQueue = append(nextQueue, &NodeIdx{nodeIdx.treeNode.Right, 2*nodeIdx.Idx + 2})
			}
		}
		currentQueue = nextQueue
	}
	return maxWidth
}

// @lc code=end

