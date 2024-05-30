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

	// 找出链表中点
	slowNode := head
	for fastNode := head; fastNode != nil && fastNode.Next != nil; {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}

	middelNode := slowNode.Next
	slowNode.Next = nil

	rightNode := reverseList(middelNode)

	// 依次合并两个链表
	// 1 2 3
	// 5 4
	for leftNode := head; leftNode != nil && rightNode != nil; {
		leftNextNode := leftNode.Next
		rightNextNode := rightNode.Next

		leftNode.Next = rightNode
		leftNode = leftNextNode

		rightNode.Next = leftNode
		rightNode = rightNextNode
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prevNode *ListNode
	for scanNode := head; scanNode != nil; {
		nextScanNode := scanNode.Next

		scanNode.Next = prevNode

		prevNode = scanNode

		scanNode = nextScanNode
	}

	return prevNode
}

// @lc code=end

