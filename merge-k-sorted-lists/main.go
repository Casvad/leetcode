package main

//import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//[[4,5],[3,4],[2,6]]
// headNode 1 -> 1
// currentNode 1 -> 1

// minimumNode 1
// minimumIndex 1
// continueProcess true
func mergeKLists(lists []*ListNode) *ListNode {

	var headNode *ListNode
	var currentNode *ListNode

	for {
		var minimumNode *ListNode
		var minimumIndex int
		continueProcess := false
		for i := range lists {
			if lists[i] != nil && (minimumNode == nil || lists[i].Val < minimumNode.Val) {
				continueProcess = true
				minimumNode = lists[i]
				minimumIndex = i
			}
		}
		if !continueProcess {
			break
		}
		if headNode == nil {
			headNode = minimumNode
			currentNode = minimumNode
		} else {
			currentNode.Next = minimumNode
			currentNode = minimumNode
		}

		lists[minimumIndex] = lists[minimumIndex].Next

	}

	return headNode
}
