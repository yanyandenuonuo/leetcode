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
	// solution1: 递归
	/**
	if len(nums) == 0 {
		return nil
	}

	maxValIdx := 0
	for idx := 1; idx < len(nums); idx += 1 {
		if nums[idx] > nums[maxValIdx] {
			maxValIdx = idx
		}
	}

	return &TreeNode{
		Val:   nums[maxValIdx],
		Left:  constructMaximumBinaryTree(nums[:maxValIdx]),
		Right: constructMaximumBinaryTree(nums[maxValIdx+1:]),
	}
	*/

	// solution2: 单调栈
	if len(nums) == 0 {
		return nil
	}

	leftMax, rightMax := make([]int, len(nums)), make([]int, len(nums))

	numStack := []int{-1}

	treeNodeList := make([]*TreeNode, len(nums))

	for idx := range nums {
		// 设置右侧最大值的默认值为-1
		rightMax[idx] = -1

		for numStack[len(numStack)-1] >= 0 && nums[idx] > nums[numStack[len(numStack)-1]] {
			rightMax[numStack[len(numStack)-1]] = idx
			numStack = numStack[:len(numStack)-1]
		}

		leftMax[idx] = numStack[len(numStack)-1]
		numStack = append(numStack, idx)

		treeNodeList[idx] = &TreeNode{
			Val:   nums[idx],
			Left:  nil,
			Right: nil,
		}
	}

	var rootNode *TreeNode
	for idx, treeNode := range treeNodeList {
		if leftMax[idx] == -1 && rightMax[idx] == -1 {
			rootNode = treeNode
		} else if rightMax[idx] == -1 || (leftMax[idx] != -1 && nums[leftMax[idx]] < nums[rightMax[idx]]) {
			// 较小侧节点的相反方向子节点
			treeNodeList[leftMax[idx]].Right = treeNodeList[idx]
		} else {
			treeNodeList[rightMax[idx]].Left = treeNodeList[idx]
		}
	}

	return rootNode
}

// @lc code=end

