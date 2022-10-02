/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	headNode := new(ListNode)
	cursorNode := headNode
	for {
		if list1 == nil && list2 != nil {
			cursorNode.Next = list2
			break
		}
		if list1 != nil && list2 == nil {
			cursorNode.Next = list1
			break
		}
		if list1 == nil && list2 == nil {
			break
		}
		if list1.Val < list2.Val {
			cursorNode.Next = list1
			cursorNode = cursorNode.Next
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			cursorNode.Next = list2
			cursorNode = cursorNode.Next
			list2 = list2.Next
		} else {
			cursorNode.Next = list1
			cursorNode = cursorNode.Next
			list1 = list1.Next
			cursorNode.Next = list2
			cursorNode = cursorNode.Next
			list2 = list2.Next
		}
	}
	return headNode.Next
}

// @lc code=end

