/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	// 				5
	//			   /  \
	//			1		8
	//				  /   \
	//				3	   11
	//
	// solution1: 如下checkTreeNode实现
	// check 5 math.min math.max
	// check 1 math.min 5 && check 8 5 math.max
	// check 3 5 5 && check 11 8 math.max

	// solution2: 中序遍历，如为二叉搜索树则得到的数组一定为升序。[1, 5, 3, 8, 11]
	return checkTreeNode(root, math.MinInt64, math.MaxInt64)
}

func checkTreeNode(root *TreeNode, minValue, maxValue int) bool {
	if root == nil {
		return true
	}

	if root.Val <= minValue || root.Val >= maxValue {
		return false
	}

	return checkTreeNode(root.Left, minValue, root.Val) && checkTreeNode(root.Right, root.Val, maxValue)
}

// @lc code=end

