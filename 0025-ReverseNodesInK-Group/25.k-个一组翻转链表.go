/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	var prePtr, curPtr *ListNode
	findK, reverseCount := false, 0
	for scanPtr := head; scanPtr != nil; scanPtr = scanPtr.Next {
		if !findK {
			nodeCount := 0
			// 查看是否存在k个节点
			for findPtr := scanPtr; findPtr != nil; findPtr = findPtr.Next {
				nodeCount += 1
				if nodeCount == k {
					findK = true
					break
				}
			}

			// 不存在k个节点
			if !findK {
				break
			}
		}
		if curPtr == nil {
			curPtr = scanPtr
		} else {
			curPtr.Next = scanPtr.Next
			if prePtr != nil {
				scanPtr.Next = prePtr.Next
				prePtr.Next = scanPtr
			} else {
				scanPtr.Next = head
				head = scanPtr
			}
			scanPtr = curPtr
		}

		reverseCount += 1
		if reverseCount == k {
			findK = false
			reverseCount = 0

			prePtr = curPtr
			curPtr = nil
		}
	}
	return head
}

// @lc code=end

