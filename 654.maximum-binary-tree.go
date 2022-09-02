/*
 * @lc app=leetcode id=654 lang=golang
 *
 * [654] Maximum Binary Tree
 *
 * https://leetcode.com/problems/maximum-binary-tree/description/
 *
 * algorithms
 * Medium (80.36%)
 * Likes:    2064
 * Dislikes: 247
 * Total Accepted:    142.6K
 * Total Submissions: 177.5K
 * Testcase Example:  '[3,2,1,6,0,5]'
 *
 * 
 * Given an integer array with no duplicates. A maximum tree building on this
 * array is defined as follow:
 * 
 * The root is the maximum number in the array. 
 * The left subtree is the maximum tree constructed from left part subarray
 * divided by the maximum number.
 * The right subtree is the maximum tree constructed from right part subarray
 * divided by the maximum number. 
 * 
 * 
 * 
 * 
 * Construct the maximum tree by the given array and output the root node of
 * this tree.
 * 
 * 
 * Example 1:
 * 
 * Input: [3,2,1,6,0,5]
 * Output: return the tree root node representing the following tree:
 * 
 * ⁠     6
 * ⁠   /   \
 * ⁠  3     5
 * ⁠   \    / 
 * ⁠    2  0   
 * ⁠      \
 * ⁠       1
 * 
 * 
 * 
 * Note:
 * 
 * The size of the given array will be in the range [1,1000].
 * 
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return buildNode(nums)
}


func buildNode(subList []int) *TreeNode {
    if len(subList) == 0 {
		return nil
	}
	maxVal := subList[0]
	splitIdx := 0
	for idx := 1; idx < len(subList); idx += 1 {
		if maxVal < subList[idx] {
			maxVal = subList[idx]
			splitIdx = idx
		}
	}
	root := new(TreeNode)
	root.Val = maxVal
	if splitIdx > 0 {
		root.Left = buildNode(subList[:splitIdx])
	} else {
		root.Left = nil
	}

	if splitIdx < len(subList)-1 {
		root.Right = buildNode(subList[splitIdx+1:])
	} else {
		root.Right = nil
	}
	return root
}
// @lc code=end

