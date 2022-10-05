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
	//	solution1: hashmap
	//  solution2: length diff: 先求2个链表长度可以得到差值diff，让较长的链表指针先走diff步逐个比较即可
	//  solution3: 走2个指针走相同步数：走到指针末尾后指向另一个指针的头部，则A指针走了(LA-L3) + (C1-C2) + (LB1-LB2)，B指针走了(LB1-LB2) + (C1-C2) + (LA-L3)
	//		LA1	->	LA2	->	LA3
	//							->	C1	->	C2
	//				LB1	->	LB2
	if headA == nil || headB == nil {
		return nil
	}
	pointA := headA
	pointB := headB

	for pointA != pointB {
		if pointA != nil {
			pointA = pointA.Next
		} else {
			pointA = headB
		}

		if pointB != nil {
			pointB = pointB.Next
		} else {
			pointB = headA
		}
	}
	return pointA
}

// @lc code=end

