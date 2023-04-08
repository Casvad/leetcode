package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	x := 5 / 10
	fmt.Printf("%d\n", x)

	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}

	l3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l4 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	runExample(l3, l4)
	runExample(l1, l2)
}

func runExample(l1 *ListNode, l2 *ListNode) {
	result := addTwoNumbers(l1, l2)
	fmt.Printf("The result of %v is %d\n", l1, result)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sumNode := ListNode{}

	addTwoNumbersSum(l1, l2, &sumNode, 0)

	return &sumNode
}

func addTwoNumbersSum(l1 *ListNode, l2 *ListNode, l3 *ListNode, additionalSum int) {

	if l1 != nil && l2 != nil {
		l3.Val = l1.Val + l2.Val + additionalSum
	} else if l1 != nil {
		l3.Val = l1.Val + additionalSum
	} else if l2 != nil {
		l3.Val = l2.Val + additionalSum
	}
	var l1Next *ListNode
	var l2Next *ListNode
	if l1 != nil {
		l1Next = l1.Next
	}
	if l2 != nil {
		l2Next = l2.Next
	}

	toAdd := 0
	if l3.Val > 9 {
		toAdd = l3.Val / 10
		l3.Val = l3.Val % 10
	}
	if l1Next == nil && l2Next == nil {
		//additionalSum += toAdd
		newNodes := make([]ListNode, 0)

		for toAdd > 0 {
			digit := toAdd % 10
			newNodes = append(newNodes, ListNode{Val: digit})
			toAdd /= 10
		}
		if len(newNodes) > 0 {
			l3.Next = &newNodes[0]
		}
		for i := 0; i < len(newNodes)-1; i++ {
			newNodes[i].Next = &newNodes[i+1]
		}
		return
	}

	nextSum := ListNode{}
	l3.Next = &nextSum

	addTwoNumbersSum(l1Next, l2Next, &nextSum, toAdd)
}
