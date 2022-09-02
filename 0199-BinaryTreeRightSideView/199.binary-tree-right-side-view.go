/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 *
 * https://leetcode.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (54.84%)
 * Likes:    2944
 * Dislikes: 165
 * Total Accepted:    350.2K
 * Total Submissions: 634.4K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * Given a binary tree, imagine yourself standing on the right side of it,
 * return the values of the nodes you can see ordered from top to bottom.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,3,null,5,null,4]
 * Output: [1, 3, 4]
 * Explanation:
 * 
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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
 package leetcode
 
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0, 32)
	queueList := make([]*TreeNode, 0, 32)
	queueList = append(queueList, root)

	currentLen := 1
	nexLen := 0
	for idx := 0; idx < currentLen; idx += 1 {
		if queueList[idx].Left != nil {
			queueList = append(queueList, queueList[idx].Left)
			nexLen += 1
		}
		if queueList[idx].Right != nil {
			queueList = append(queueList, queueList[idx].Right)
			nexLen += 1
		}

		if currentLen-idx == 1 {
			res = append(res, queueList[idx].Val)
			queueList = queueList[currentLen:]
			idx = -1
			currentLen = nexLen
			nexLen = 0
			if currentLen == 0 {
				break
			}
		}

	}
	return res
}
// @lc code=end

