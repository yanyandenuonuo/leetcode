/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pointA := headA
	pointB := headB

	for {
		if pointA == pointB {
			return pointA
		}

		if pointA.Next == nil && pointB.Next == nil {
			return nil
		}

		if pointA.Next != nil {
			pointA = pointA.Next
		} else if pointB.Next != nil {
			pointA = headB
		}

		if pointB.Next != nil {
			pointB = pointB.Next
		} else if pointA.Next != nil {
			pointB = headA
		}
	}
}

// @lc code=end

