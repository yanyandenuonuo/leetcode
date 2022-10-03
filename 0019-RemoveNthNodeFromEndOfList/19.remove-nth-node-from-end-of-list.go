/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slowNode, fastNode := head, head

	// fast run n step
	for idx := 0; idx < n; idx += 1 {
		if fastNode != nil {
			fastNode = fastNode.Next
		}
	}

	for {
		if fastNode == nil {
			// link is shorter than n
			return head.Next
		} else if fastNode.Next == nil {
			// fast is last
			slowNode.Next = slowNode.Next.Next
			return head
		}
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}
}

// @lc code=end

