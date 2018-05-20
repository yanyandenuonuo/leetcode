/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        Integer leftValue = 0;
        Integer posValue = 0;
        Integer sumValue = 0;
        ListNode response = new ListNode(0);
        ListNode currentNode = response;
        while(true) {
            sumValue = l1.val + l2.val + leftValue;
            posValue = sumValue % 10;
            leftValue = sumValue / 10;
            currentNode.next = new ListNode(posValue);
            currentNode = currentNode.next;
            if(l1.next == null && l2.next == null) {
                break;
            }
            if(l1.next != null) {
                l1 = l1.next;
            } else {
                if(leftValue == 0) {
                    currentNode.next = l2.next;
                    break;
                }else {
                    l1 = new ListNode(0);
                }
            }
            if(l2.next != null){
                l2 = l2.next;
            } else {
                if(leftValue == 0) {
                    currentNode.next = l1;
                    break;
                }else {
                    l2 = new ListNode(0);
                }
            }
        }
        if(leftValue > 0) {
            currentNode.next = new ListNode(leftValue);
        }
        return response.next;
    }
}
