package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addNodes(nil, l1, l2, 0)
}

func addNodes(prev *ListNode, l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	v1, n1 := value(l1)
	v2, n2 := value(l2)
	v := v1 + v2 + carry
	c := 0
	if v > 9 {
		v -= 10
		c = 1
	}

	node := &ListNode{Val: v}
	if prev != nil {
		prev.Next = node
	}

	if n1 != nil || n2 != nil || c != 0 {
		addNodes(node, n1, n2, c)
	}
	return node
}

func value(n *ListNode) (int, *ListNode) {
	if n == nil {
		return 0, nil
	}
	return n.Val, n.Next
}
