/*
 * @lc app=leetcode.cn id=2095 lang=golang
 *
 * [2095] 删除链表的中间节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	preIdx, slowIdx, fastIdx := head, head, head
	for {
		if fastIdx == nil || fastIdx.Next == nil {
			if preIdx == slowIdx {
				return nil
			}
			preIdx.Next = slowIdx.Next
			return head
		}
		fastIdx = fastIdx.Next.Next
		preIdx = slowIdx
		slowIdx = slowIdx.Next
	}
}

// @lc code=end

