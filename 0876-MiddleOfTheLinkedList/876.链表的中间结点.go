/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
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
	// solution: 快慢指针
	slowIdx := head
	for fastIdx := head; fastIdx != nil && fastIdx.Next != nil; {
		slowIdx = slowIdx.Next
		fastIdx = fastIdx.Next.Next
	}

	return slowIdx
}

// @lc code=end

