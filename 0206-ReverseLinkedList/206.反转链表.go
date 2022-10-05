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
	// solution1: 构建新链表
	// if head == nil {
	// 	return nil
	// }
	// var newHead *ListNode
	// pointNode := head
	// for pointNode != nil {
	// 	newHead = &ListNode{
	// 		Val:  pointNode.Val,
	// 		Next: newHead,
	// 	}

	// 	pointNode = pointNode.Next
	// }
	// return newHead

	// solution2: 原地更改指向
	if head == nil {
		return nil
	}
	point := head.Next
	head.Next = nil

	for point != nil {
		nextPoint := point.Next
		point.Next = head
		head = point
		point = nextPoint
	}
	return head
}

// @lc code=end

