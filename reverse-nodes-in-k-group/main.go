package main

import "fmt"

func main() {

	node := []int{1, 2, 3, 4, 5}

	x := reverseKGroup(intArray(node).toNode(), 2)

	print(x.ToString())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type intArray []int

func (t intArray) toNode() *ListNode {
	head := &ListNode{Val: t[0], Next: nil}
	currentNode := head
	for i := 1; i < len(t); i++ {
		newNode := &ListNode{Val: t[i], Next: nil}
		currentNode.Next = newNode
		currentNode = newNode
	}

	return head
}
func (t *ListNode) ToString() string {
	if t == nil {
		return "\n"
	} else {
		return fmt.Sprintf("%v -> %v", t.Val, t.Next.ToString())
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	var newHead *ListNode
	stack := make([]*ListNode, 0)
	iteratorNode := head
	var previousBatchLastNode *ListNode
	for iteratorNode != nil {
		nextNode := iteratorNode.Next
		stack = append(stack, iteratorNode)
		if len(stack) == k {
			stack[0].Next = iteratorNode.Next
			for i := len(stack) - 1; i > 0; i-- {
				stack[i].Next = stack[i-1]
			}
			if newHead == nil {
				newHead = iteratorNode
			}
			if previousBatchLastNode != nil {
				previousBatchLastNode.Next = iteratorNode
			}
			previousBatchLastNode = stack[0]
			stack = make([]*ListNode, 0)
		}
		iteratorNode = nextNode
	}

	return newHead
}
