/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	resNode := new(ListNode)
	resNode.Next = &ListNode{
		Val:  head.Val,
		Next: nil,
	}

	for scanNode := head.Next; scanNode != nil; scanNode = scanNode.Next {
		preScanInsert := resNode
		insertSuccess := false
		for scanInsert := resNode.Next; scanInsert != nil; scanInsert = scanInsert.Next {
			if scanNode.Val <= scanInsert.Val {
				preScanInsert.Next = &ListNode{
					Val:  scanNode.Val,
					Next: scanInsert,
				}
				insertSuccess = true
				break
			}
			preScanInsert = preScanInsert.Next
		}
		if !insertSuccess {
			preScanInsert.Next = &ListNode{
				Val:  scanNode.Val,
				Next: nil,
			}
		}
	}

	return resNode.Next
}

// @lc code=end

