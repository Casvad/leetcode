package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}

	list2CurrentNode := list2
	list1CurrentNode := list1

	for {

		if list2CurrentNode == nil {
			break
		}
		if list1CurrentNode.Next == nil {
			list1CurrentNode.Next = list2CurrentNode
			break
		}

		if list1CurrentNode.Next.Val > list2CurrentNode.Val {
			tmp := list2CurrentNode
			list2CurrentNode = list2CurrentNode.Next
			tmp.Next = list1CurrentNode.Next
			list1CurrentNode.Next = tmp
		} else {
			list1CurrentNode = list1CurrentNode.Next
		}

	}

	return list1
}
