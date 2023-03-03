/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var preLeftPtr, leftPtr, rightPtr *ListNode
	for idx, scanPtr := 1, head; scanPtr != nil; scanPtr = scanPtr.Next {
		if idx == left {
			leftPtr = scanPtr
		}

		if idx == right {
			rightPtr = scanPtr
		}

		if leftPtr == nil {
			preLeftPtr = scanPtr
		} else if leftPtr == rightPtr {
			break
		}

		// 反转状态
		if leftPtr != nil && scanPtr != leftPtr {
			leftPtr.Next = scanPtr.Next
			if preLeftPtr != nil {
				scanPtr.Next = preLeftPtr.Next
				preLeftPtr.Next = scanPtr
			} else {
				scanPtr.Next = head
				head = scanPtr
			}

			scanPtr = leftPtr
		}

		if idx == right {
			break
		}

		idx += 1
	}
	return head
}

// @lc code=end

