package main

// Total nodes = a + b, where a is the number of nodes before cycle entrance and b is the number of nodes in cycle
// When there is cycle and fast meets slow, slow moved s steps and fast moved 2s steps
// fast = 2 * slow, fast = slow + nb, where n is the number of times slow cycled
// So we get fast = 2nb, slow = nb
// We can reach to cycle entrance by k = a + nb steps
// So we just need to move slow another a steps to find the cycle entrance
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	for fast, slow := head, head; fast != nil && fast.Next != nil; {
		fast, slow = fast.Next.Next, slow.Next
		// has cycle
		if fast == slow {
			// Reset fast to head, both fast and slow move a steps can reach to cycle entrance
			fast = head
			for fast != slow {
				fast, slow = fast.Next, slow.Next
			}
			return slow
		}
	}
	return nil
}
