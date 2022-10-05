/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// solution1: 使用栈存储两个链表数据
	// solution2: 反转链表

	point1 := l1.Next
	l1.Next = nil

	// 反转链表1
	for point1 != nil {
		nextPoint := point1.Next
		point1.Next = l1
		l1 = point1
		point1 = nextPoint
	}
	point1 = l1

	point2 := l2.Next
	l2.Next = nil

	// 反转链表1
	for point2 != nil {
		nextPoint := point2.Next
		point2.Next = l2
		l2 = point2
		point2 = nextPoint
	}
	point2 = l2

	var pointSum *ListNode
	carry := 0
	for point1 != nil && point2 != nil {
		bitSum := point1.Val + point2.Val + carry
		carry = bitSum / 10
		currentNode := &ListNode{
			Val:  (bitSum) % 10,
			Next: pointSum,
		}
		pointSum = currentNode

		point1 = point1.Next
		point2 = point2.Next
	}

	for point1 != nil {
		bitSum := point1.Val + carry
		carry = bitSum / 10

		currentNode := &ListNode{
			Val:  (bitSum) % 10,
			Next: pointSum,
		}
		pointSum = currentNode

		point1 = point1.Next
	}

	for point2 != nil {
		bitSum := point2.Val + carry
		carry = bitSum / 10

		currentNode := &ListNode{
			Val:  (bitSum) % 10,
			Next: pointSum,
		}
		pointSum = currentNode

		point2 = point2.Next
	}

	if carry > 0 {
		currentNode := &ListNode{
			Val:  carry % 10,
			Next: pointSum,
		}
		pointSum = currentNode
	}

	return pointSum
}

// @lc code=end

