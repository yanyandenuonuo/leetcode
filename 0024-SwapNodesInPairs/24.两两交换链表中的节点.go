/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// solution1: 递归
	// solution2: 遍历
	preNode := &ListNode{
		Val:  0,
		Next: head,
	}

	for nextNode := preNode; nextNode != nil && nextNode.Next != nil; {
		firstNode := nextNode.Next
		secondNode := firstNode.Next
		if secondNode == nil {
			break
		}

		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		nextNode.Next = secondNode
		nextNode = firstNode
	}

	return preNode.Next
}

// @lc code=end

