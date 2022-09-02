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
	slowNode := head
	fastNode := head

	// fast run n step
	for idx := 0; idx < n; idx += 1 {
		if fastNode != nil {
			fastNode = fastNode.Next
		}
	}

	for {
		if fastNode == nil {
			// fast is nil
			return head.Next
		} else if fastNode.Next == nil {
			// fast is last
			if slowNode.Next != nil {
				slowNode.Next = slowNode.Next.Next
			} else {
				return nil
			}
			return head
		}
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}
}
// @lc code=end

