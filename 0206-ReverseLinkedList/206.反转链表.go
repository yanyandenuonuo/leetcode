/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	pointNode := head
	for pointNode != nil {
		newHead = &ListNode{
			Val:  pointNode.Val,
			Next: newHead,
		}

		pointNode = pointNode.Next
	}
	return newHead
}

// @lc code=end

