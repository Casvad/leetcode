package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	tmp := head
	prev := head
	var currentHead *ListNode
	for tmp != nil {

		toSwap := tmp.Next
		if toSwap == nil {
			break
		}
		tmp.Next = toSwap.Next
		toSwap.Next = tmp

		if currentHead == nil {
			currentHead = toSwap
		} else {
			prev.Next = toSwap
		}
		prev = tmp
		tmp = tmp.Next
	}

	if currentHead == nil {
		return head
	}

	return currentHead
}
