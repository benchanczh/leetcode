package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur, count := head, 0
	// Check if there are enough node for reverse, return head if not
	for count < k {
		// base case
		if cur == nil {
			return head
		}
		cur = cur.Next
		count++
	}

	prev := reverseKGroup(cur, k)

	// Reverse one by one in a group
	for count > 0 {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
		count--
	}

	return prev
}
