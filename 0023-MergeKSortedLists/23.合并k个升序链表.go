/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// solution1: 顺序合并
	//			  依次将k个链表与合并后的链表进行合并mergeTwo
	// solution2: 分治合并
	//			  依次将k个链表中的两两合并，再对合并后的(k+1)/2个链表进行两两合并
	// solution3: 利用小顶堆
	//			  使用k个ListNode的第一个元素构成一个长度为k的小顶堆
	//			  依次从堆顶取出元素并放入该元素Next元素
	//			  再对堆进行排序
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	}

	// val, idx
	sortIdx := make([][]int, 0, len(lists))
	for idx, node := range lists {
		if node == nil {
			continue
		}
		sortIdx = append(sortIdx, []int{node.Val, idx})
	}

	if len(sortIdx) == 0 {
		return nil
	}

	sortIdx = heapSort(sortIdx)

	sortList := new(ListNode)
	head := sortList
	minListIdx := sortIdx[0][1]
	sortList.Val = lists[minListIdx].Val
	if lists[minListIdx].Next != nil {
		lists[minListIdx] = lists[minListIdx].Next
		sortIdx[0] = []int{lists[minListIdx].Val, minListIdx}
		heapSort(sortIdx)
	} else {
		sortIdx = sortIdx[1:]
	}

	for len(sortIdx) != 0 {
		minListIdx = sortIdx[0][1]
		sortList.Next = &ListNode{
			Val: lists[minListIdx].Val,
		}
		sortList = sortList.Next

		if lists[minListIdx].Next != nil {
			lists[minListIdx] = lists[minListIdx].Next
			sortIdx[0] = []int{lists[minListIdx].Val, minListIdx}
			heapSort(sortIdx) // 这里其实走插值排序更快log(n)
		} else {
			sortIdx = sortIdx[1:]
		}
	}

	return head
}

func heapSort(nums [][]int) [][]int {
	// 构建小顶堆
	for idx := len(nums)/2 - 1; idx >= 0; idx -= 1 {
		heapify(nums, idx, len(nums))
	}

	for idx := len(nums) - 1; idx > 0; idx -= 1 {
		nums[idx], nums[0] = nums[0], nums[idx]
		heapify(nums, 0, idx)
	}
	return nums
}

func heapify(nums [][]int, idx, length int) {
	leftChildIdx := 2*idx + 1
	rightChildIdx := 2*idx + 2
	largestIdx := idx
	if leftChildIdx < length && nums[leftChildIdx][0] > nums[largestIdx][0] {
		largestIdx = leftChildIdx
	}

	if rightChildIdx < length && nums[rightChildIdx][0] > nums[largestIdx][0] {
		largestIdx = rightChildIdx
	}
	if largestIdx != idx {
		nums[largestIdx], nums[idx] = nums[idx], nums[largestIdx]
		heapify(nums, largestIdx, length)
	}
}

// @lc code=end

