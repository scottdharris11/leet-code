package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// Return if nil
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Compute length of node list
	var pEnd *ListNode
	l := 1
	for pEnd = head; pEnd.Next != nil; pEnd = pEnd.Next {
		l++
	}

	// establish index of the new ending node
	// account for offset being greater than node list length
	work := l - k
	if k > l {
		work = k % l
		if work > 0 {
			work = l - work
		}
	}

	// If wrapping completely around...just return
	if work == 0 {
		return head
	}

	// identify the new end and beginning nodes
	var nBegin *ListNode
	node := head
	for {
		work--
		if work == 0 {
			nBegin = node.Next
			pEnd.Next = head
			node.Next = nil
			break
		}
		node = node.Next
	}
	return nBegin
}
