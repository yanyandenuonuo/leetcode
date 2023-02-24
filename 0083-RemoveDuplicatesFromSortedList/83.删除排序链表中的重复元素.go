/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	valMap := make(map[int]bool)
	for currentPtr, prePtr := head, head; currentPtr != nil; currentPtr = currentPtr.Next {
		if !valMap[currentPtr.Val] {
			valMap[currentPtr.Val] = true
			prePtr = currentPtr
			continue
		}
		prePtr.Next = currentPtr.Next
	}
	return head
}

// @lc code=end

