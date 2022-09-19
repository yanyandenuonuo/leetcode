/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	// 前序遍历: root, root.left, root.right
	// 中序遍历: root.left, root, root.right
	// 后续遍历: root.left, root.right, root
	// 总结：前、中、后是说的root所在的位置
	return getNodePreorderArray(root)
}

func getNodePreorderArray(curNode *Node) []int {
	res := make([]int, 0)
	if curNode == nil {
		return res
	}
	res = append(res, curNode.Val)
	for _, childNode := range curNode.Children {
		res = append(res, getNodePreorderArray(childNode)...)
	}
	return res
}

// @lc code=end

