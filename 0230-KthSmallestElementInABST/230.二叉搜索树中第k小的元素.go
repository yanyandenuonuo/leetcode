/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	// solution1: 中序遍历，然后返回数组中第k-1个值
	// solution2: 统计节点数字
	nodeCountMap := make(map[*TreeNode]int, k*2)
	for {
		nodeCount := findNodeCount(nodeCountMap, root.Left)
		if nodeCount < k-1 {
			root = root.Right
			k -= nodeCount + 1
		} else if nodeCount > k-1 {
			root = root.Left
		} else {
			return root.Val
		}
	}
}

func findNodeCount(nodeCountMap map[*TreeNode]int, root *TreeNode) int {
	if root == nil {
		return 0
	}

	if _, isExist := nodeCountMap[root]; isExist {
		return nodeCountMap[root]
	}

	nodeCountMap[root] = 1 + findNodeCount(nodeCountMap, root.Left) + findNodeCount(nodeCountMap, root.Right)
	return nodeCountMap[root]
}

// @lc code=end

