/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
 *
 * https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (55.79%)
 * Likes:    1659
 * Dislikes: 144
 * Total Accepted:    162.3K
 * Total Submissions: 290.8K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * Given the root of a Binary Search Tree and a target number k, return true if
 * there exist two elements in the BST such that their sum is equal to the
 * given target.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [5,3,6,2,4,null,7], k = 9
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root = [5,3,6,2,4,null,7], k = 28
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: root = [2,1,3], k = 4
 * Output: true
 *
 *
 * Example 4:
 *
 *
 * Input: root = [2,1,3], k = 1
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: root = [2,1,3], k = 3
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 10^4].
 * -10^4 <= Node.val <= 10^4
 * root is guaranteed to be a valid binary search tree.
 * -10^5 <= k <= 10^5
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
func findTarget(root *TreeNode, k int) bool {
	// solution1: 中序遍历转数组再进行2sum查找
	// solution2: DFS
	/**
	if root == nil {
		return false
	}

	return findNode(root, root, k)
	*/

	// solution3: DFS + hash表
	if root == nil {
		return false
	}

	valMap := make(map[int]bool)

	var findVal func(treeNode *TreeNode) bool
	findVal = func(treeNode *TreeNode) bool {
		if treeNode == nil {
			return false
		}

		if valMap[k-treeNode.Val] {
			return true
		}

		valMap[treeNode.Val] = true

		return findVal(treeNode.Left) || findVal(treeNode.Right)
	}
	return findVal(root)

	// solution4: DFS + 中序遍历 + 双指针
	// solution5: 迭代 + 中序遍历 + 双指针
}

func findNode(root *TreeNode, searchNode *TreeNode, k int) bool {
	if searchNode == nil {
		return false
	}

	if findNum(root, searchNode, k-searchNode.Val) {
		return true
	}

	if findNode(root, searchNode.Left, k) {
		return true
	}

	if findNode(root, searchNode.Right, k) {
		return true
	}

	return false
}

func findNum(root *TreeNode, searchNode *TreeNode, num int) bool {
	if root == nil {
		return false
	}

	if root.Val > num {
		return findNum(root.Left, searchNode, num)
	} else if root.Val < num {
		return findNum(root.Right, searchNode, num)
	} else if root.Val == num && root != searchNode {
		return true
	}

	return false
}

// @lc code=end

