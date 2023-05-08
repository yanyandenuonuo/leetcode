/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	// solution1: 前序遍历获得数组再构造链表
	// solution2: 前序遍历+单调栈同时构造链表
	/**
	if root == nil {
		return
	}

	var preNode *TreeNode

	nodeStack := make([]*TreeNode, 0)
	nodeStack = append(nodeStack, root)

	for len(nodeStack) > 0 {
		currentNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		if preNode != nil {
			preNode.Left = nil
			preNode.Right = currentNode
		}

		// 因为栈先进后出，所以先压左边节点
		if currentNode.Right != nil {
			nodeStack = append(nodeStack, currentNode.Right)
		}

		if currentNode.Left != nil {
			nodeStack = append(nodeStack, currentNode.Left)
		}

		preNode = currentNode
	}
	*/

	// solution3: 寻找链表前节点，本质是将节点的右子节点转移到左子节点的最右叶子节点的右节点，然后交换当前节点的左右节点
	for currentNode := root; currentNode != nil; currentNode = currentNode.Right {
		if currentNode.Left == nil {
			continue
		}

		nextNode := currentNode.Left
		rightPreNode := nextNode
		for rightPreNode.Right != nil {
			rightPreNode = rightPreNode.Right
		}

		// 将当前节点的右节点链接到当前左节点的最右叶子节点
		rightPreNode.Right = currentNode.Right

		// 交换左右节点
		currentNode.Right = currentNode.Left
		currentNode.Left = nil
	}
}

// @lc code=end

