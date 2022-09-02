/*
 * @lc app=leetcode id=876 lang=golang
 *
 * [876] Middle of the Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slowNode := head
	fastNode := head
	for {
		if fastNode == nil || fastNode.Next == nil {
			return slowNode
		}
		slowNode = slowNode.Next
		fastNode = fastNode.Next
		if fastNode != nil {
			fastNode = fastNode.Next
		}
	}
}
// @lc code=end

