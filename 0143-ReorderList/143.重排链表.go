/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	// solution1: 转化为数组再按要求构建新链表
	// solution2: 将链表切割为前后两部分，对后面一部分进行逆序，与前面一部分链表合并即可
	if head == nil || head.Next == nil {
		return
	}
	preMidPoint, slowPoint := head, head
	for fastPoint := head; fastPoint != nil; fastPoint = fastPoint.Next {
		preMidPoint = slowPoint
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next
		if fastPoint == nil {
			break
		}
	}
	preMidPoint.Next = nil
	list2 := reverseList(slowPoint)

	res := new(ListNode)
	// 依次合并两个链表
	// 1 2 3
	// 5 4
	for list1Node, list2Node, tailNode := head, list2, res; list1Node != nil || list2Node != nil; {
		tailNode.Next = list1Node
		tailNode = tailNode.Next

		list1Node = list1Node.Next

		if list2Node == nil {
			break
		}

		tailNode.Next = list2Node
		tailNode = tailNode.Next

		list2Node = list2Node.Next
	}
	head = res.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	for scanNode := head; scanNode != nil; {
		currentScanNode := scanNode
		scanNode = scanNode.Next

		currentScanNode.Next = newHead
		newHead = currentScanNode
	}
	return newHead
}

// @lc code=end

