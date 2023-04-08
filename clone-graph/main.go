package main

import (
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {

	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 2}
	node4 := &Node{Val: 2}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node2}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	runExample(node1)
}

func runExample(input *Node) {
	result := cloneGraph(input)
	fmt.Printf("The result of %v is %d\n", input, result)
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	cloneMap := map[int]*Node{}
	return cloneGraphWithMap(node, &cloneMap)
}

func cloneGraphWithMap(node *Node, nodeMap *map[int]*Node) *Node {

	nodeMapElement, found := (*nodeMap)[node.Val]
	if found {
		return nodeMapElement
	}
	newNode := Node{}
	newNeighbors := make([]*Node, len(node.Neighbors))
	(*nodeMap)[node.Val] = &newNode

	for i, neighbor := range node.Neighbors {
		newNeighbor := cloneGraphWithMap(neighbor, nodeMap)
		newNeighbors[i] = newNeighbor
	}

	newNode.Val = node.Val
	newNode.Neighbors = newNeighbors

	return &newNode
}
