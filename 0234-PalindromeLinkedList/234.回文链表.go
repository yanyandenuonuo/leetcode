/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// solution1: 转为数组再使用双指针判断
	// solution2: 使用快慢指针找出链表中点，翻转后半部链表，然后与前半部分链表比对
	// solution3: 利用递归的特性，将左边节点与右边节点相比较
	leftNode := head
	var checkNode func(currentNode *ListNode) bool
	checkNode = func(currentNode *ListNode) bool {
		if currentNode == nil {
			return true
		}

		if !checkNode(currentNode.Next) {
			return false
		}

		if currentNode.Val != leftNode.Val {
			return false
		}

		leftNode = leftNode.Next

		return true
	}

	return checkNode(head)
}

// @lc code=end

