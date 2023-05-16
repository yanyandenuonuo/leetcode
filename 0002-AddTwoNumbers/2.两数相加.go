/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
	// solution: 模拟
	point1 := l1
	point2 := l2
	pointSum := new(ListNode)
	point3 := pointSum
	carry := 0
	prePoint := new(ListNode)
	for point1 != nil && point2 != nil {
		bitSum := point1.Val + point2.Val + carry
		carry = bitSum / 10

		point3.Val = (bitSum) % 10
		point3.Next = new(ListNode)
		prePoint = point3
		point3 = point3.Next

		point1 = point1.Next
		point2 = point2.Next
	}

	for point1 != nil {
		bitSum := point1.Val + carry
		carry = bitSum / 10

		point3.Val = (bitSum) % 10
		point3.Next = new(ListNode)
		prePoint = point3
		point3 = point3.Next

		point1 = point1.Next
	}

	for point2 != nil {
		bitSum := point2.Val + carry
		carry = bitSum / 10

		point3.Val = (bitSum) % 10
		point3.Next = new(ListNode)
		prePoint = point3
		point3 = point3.Next

		point2 = point2.Next
	}

	if carry > 0 {
		point3.Val = carry
	} else {
		prePoint.Next = nil
	}

	return pointSum
}

// @lc code=end

