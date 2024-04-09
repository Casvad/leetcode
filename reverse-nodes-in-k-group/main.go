package main

func main() {

	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	x := reverseKGroup(node, 2)

	print(x)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	return reverseKGroupFromNode(head, k, true)
}

func reverseKGroupFromNode(head *ListNode, k int, loop bool) *ListNode {
	subK := k
	for k >= 0 {
		tmp := head
		prev := head
		for i := 0; i < k; i++ {
			toSwap := tmp.Next
			if toSwap == nil {
				break
			}
			tmp.Next = toSwap.Next
			toSwap.Next = tmp

			if i == 0 {
				head = toSwap
			} else {
				prev.Next = toSwap
			}
			prev = toSwap
		}
		k--
		if k < 0 && loop {
			reverseKGroupFromNode(tmp, subK, false)
			return head
		}
	}

	return head
}
