/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N 叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	// solution1: 递归
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)

	for _, node := range root.Children {
		res = append(res, postorder(node)...)
	}

	res = append(res, root.Val)

	return res

	// solution2: 栈
	// solution3: 前序遍历，然后反转结果
}

// @lc code=end

