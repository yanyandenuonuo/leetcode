/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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
	var distinctList *ListNode
	var distinctNode *ListNode

	curNode := head
	for curNode != nil {
		// not duplicate
		if curNode.Next != nil && curNode.Val == curNode.Next.Val {
			if curNode.Next.Next == nil {
				break
			} else if curNode.Next.Val == curNode.Next.Next.Val {
				curNode = curNode.Next
			} else {
				curNode = curNode.Next.Next
			}
			continue
		}
		if distinctNode == nil {
			distinctList = &ListNode{
				Val:  curNode.Val,
				Next: nil,
			}
			distinctNode = distinctList
		} else {
			distinctNode.Next = &ListNode{
				Val:  curNode.Val,
				Next: nil,
			}
			distinctNode = distinctNode.Next
		}
		curNode = curNode.Next
	}

	return distinctList
}

// @lc code=end

