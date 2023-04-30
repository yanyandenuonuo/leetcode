/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head
	for scanNode := preHead; scanNode.Next != nil; {
		if scanNode.Next.Val == val {
			scanNode.Next = scanNode.Next.Next
		} else {
			scanNode = scanNode.Next
		}
	}

	return preHead.Next
}

// @lc code=end

