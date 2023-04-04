/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var revHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = revHead
		revHead = head
		head = next
	}
	return revHead
}