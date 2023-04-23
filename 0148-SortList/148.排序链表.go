/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	// solution1: 转为数组，排序后生成链表
	// solution2: 将链表拆分成两个链表，然后合并
	/**
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var preMid *ListNode
	slowPoint := head
	for fastPoint := head; fastPoint != nil; fastPoint = fastPoint.Next {
		preMid = slowPoint
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next
		if fastPoint == nil {
			break
		}
	}

	// 将链表切割为2段
	preMid.Next = nil

	return mergeTwoLists(sortList(head), sortList(slowPoint))
	*/

	// solution3: 将链表切割为多段，每2段进行合并，然后翻倍递增每段的长度
	//			  6 | 5 | 4 | 3 | 2 | 1		每段长度为1
	//			  5   6 | 3   4 | 1   2		合并后每段长度翻倍递增为2
	//			  3   4   5   6 | 1   2		合并后每段长度翻倍递增为4
	//			  1   2   3   4   5   6
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	listLength := 0
	for scanNode := head; scanNode != nil; scanNode = scanNode.Next {
		listLength += 1
	}

	res := new(ListNode)
	res.Next = head

	for splitLength := 1; splitLength < listLength; splitLength <<= 1 {
		tail := res
		for scanNode := res.Next; scanNode != nil; {
			// 切分链表1
			list1 := scanNode
			for idx := 1; idx < splitLength && scanNode.Next != nil; idx += 1 {
				scanNode = scanNode.Next
			}
			nextScan := scanNode.Next
			scanNode.Next = nil
			scanNode = nextScan

			// 切分链表2
			list2 := scanNode
			for idx := 1; idx < splitLength && scanNode != nil; idx += 1 {
				scanNode = scanNode.Next
			}
			if scanNode != nil {
				nextScan = scanNode.Next
				scanNode.Next = nil
				scanNode = nextScan
			}

			tail.Next = mergeTwoLists(list1, list2)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}

	return res.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	headNode := new(ListNode)
	scanNode := headNode
	for list1 != nil || list2 != nil {
		if list1 == nil && list2 != nil {
			scanNode.Next = list2
			break
		}
		if list1 != nil && list2 == nil {
			scanNode.Next = list1
			break
		}

		if list1.Val <= list2.Val {
			scanNode.Next = list1
			scanNode = scanNode.Next
			list1 = list1.Next
		} else {
			scanNode.Next = list2
			scanNode = scanNode.Next
			list2 = list2.Next
		}
	}
	return headNode.Next
}

// @lc code=end

