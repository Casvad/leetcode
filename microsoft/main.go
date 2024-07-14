package main

// you can also use imports, for example:
import "fmt"
import "errors"

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func Solution(T *Tree) int {
	visitedNumbers := make(map[string]bool)
	visitedMapNodes := make(map[string]bool)
	dFS(T, visitedMapNodes, visitedNumbers)

	return len(visitedNumbers)
}

func dFS(node *Tree, visitedMapNodes, visitedNumbers map[string]bool) {

	if node == nil {
		return
	}
	key := fmt.Sprintf("%s", node)
	if _, ok := visitedMapNodes[key]; ok {
		return
	}
	fmt.
		visitedMapNodes[key] = true
	possibleSolutions, err := resolveStringNode(node, 1)
	if err != nil {
		for _, sol := range possibleSolutions {
			if _, ok := visitedNumbers[sol]; !ok {
				visitedNumbers[sol] = true
			}
		}
	}
	dFS(node.L, visitedMapNodes, visitedNumbers)
	dFS(node.R, visitedMapNodes, visitedNumbers)
}

func resolveStringNode(node *Tree, level int) ([]string, error) {
	if node == nil {
		return []string{}, errors.New("Could not create string")
	}
	if level >= 3 {
		return []string{fmt.Sprintf("%v", node.X)}, nil
	}
	ans := make([]string, 0)
	strNode := fmt.Sprintf("%v", node.X)
	l, errL := resolveStringNode(node.L, level+1)
	r, errR := resolveStringNode(node.R, level+1)

	if errL == nil {
		for _, s := range l {
			ans = append(ans, fmt.Sprintf("%v%v", strNode, s))
		}
	}
	if errR == nil {
		for _, s := range r {
			ans = append(ans, fmt.Sprintf("%v%v", strNode, s))
		}
	}

	return ans, nil
}

func main() {

}
