/*
 * @lc app=leetcode.cn id=1110 lang=golang
 *
 * [1110] 删点成林
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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	// solution1: dfs
	delValMap := make(map[int]struct{}, len(to_delete))
	for _, delVal := range to_delete {
		delValMap[delVal] = struct{}{}
	}

	treeNodeList := make([]*TreeNode, 0)

	var delNode func(*TreeNode, bool)
	delNode = func(treeNode *TreeNode, isRoot bool) {
		if treeNode == nil {
			return
		}

		if _, isExist := delValMap[treeNode.Val]; !isExist {
			if isRoot {
				treeNodeList = append(treeNodeList, treeNode)
				isRoot = false
			}
		} else {
			isRoot = true
		}

		if treeNode.Left != nil {
			delNode(treeNode.Left, isRoot)

			if _, isExist := delValMap[treeNode.Left.Val]; isExist {
				treeNode.Left = nil
			}
		}

		if treeNode.Right != nil {
			delNode(treeNode.Right, isRoot)

			if _, isExist := delValMap[treeNode.Right.Val]; isExist {
				treeNode.Right = nil
			}
		}

		return
	}

	delNode(root, true)

	return treeNodeList

	// solution2: bfs
}

// @lc code=end

