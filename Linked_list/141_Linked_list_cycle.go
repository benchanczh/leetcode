package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	for fast, slow := head, head; fast != nil && fast.Next != nil; {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
