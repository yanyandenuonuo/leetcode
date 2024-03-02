/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序搜索树
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
func increasingBST(root *TreeNode) *TreeNode {
	// solution1: 先中序遍历，再构建只有右子节点的树

	/**
	treeRoot := new(TreeNode)
	treeNode := treeRoot

	var inOrder func(node *TreeNode) []int
	inOrder = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}

		resList := inOrder(node.Left)

		resList = append(resList, node.Val)

		resList = append(resList, inOrder(node.Right)...)

		return resList
	}

	for _, val := range inOrder(root) {
		treeNode.Right = &TreeNode{
			Val: val,
		}
		treeNode = treeNode.Right
	}

	return treeRoot.Right
	*/

	// solution2: 中序遍历过程中改变节点，将当前节点改为左子节点的右子节点
	treeRoot := new(TreeNode)
	leftNode := treeRoot

	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)

		leftNode.Right = node
		node.Left = nil
		leftNode = leftNode.Right

		inOrder(node.Right)
	}

	inOrder(root)

	return treeRoot.Right
}

// @lc code=end

