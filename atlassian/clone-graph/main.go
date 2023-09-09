package main

import (
	"fmt"
)

func main() {

	n1 := Node{
		Value:     1,
		Neighbors: make([]*Node, 0),
	}

	n2 := Node{
		Value:     2,
		Neighbors: []*Node{&n1},
	}

	n3 := Node{
		Value:     3,
		Neighbors: []*Node{&n2},
	}

	n1.Neighbors = append(n1.Neighbors, &n3)

	refNodes := make(map[int]*Node, 0)

	result := clone(&n3, &refNodes)

	fmt.Printf("result is %v\n", result)
}

type Node struct {
	Value     int
	Neighbors []*Node
}

func clone(node *Node, refNodes *map[int]*Node) *Node {

	if nod, exist := (*refNodes)[node.Value]; exist {

		return nod
	}
	newNode := Node{
		Value:     node.Value,
		Neighbors: make([]*Node, 0),
	}

	(*refNodes)[node.Value] = &newNode
	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, clone(neighbor, refNodes))
	}

	return &newNode
}
