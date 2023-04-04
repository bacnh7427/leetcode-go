/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	slowPointer := head
	fastPointer := head

	flag := 0

	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if fastPointer == slowPointer {
			flag = 1
			break
		}
	}

	if flag != 1 {
		return nil
	}

	for head != slowPointer {
		head = head.Next
		slowPointer = slowPointer.Next
	}

	return head

}