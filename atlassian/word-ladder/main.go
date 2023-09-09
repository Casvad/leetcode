package main

import (
	"fmt"
)

func main() {
	result := wordLadder([]string{"hot", "dot", "dog", "lot", "log"}, "hit", "cog")

	fmt.Printf("result is %v\n", result)
}

type Node struct {
	Word  string
	Steps int
}

func wordLadder(dictionaty []string, start, end string) int {

	words := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	dictionaty = append(dictionaty, end)

	queue := make([]Node, 0)
	queue = append(queue, Node{
		Word:  start,
		Steps: 1,
	})

	for {
		if len(queue) == 0 {
			break
		}

		currentWordNode, subQueue := queue[0], queue[1:]
		queue = subQueue
		currentWord := currentWordNode.Word

		if currentWord == end {
			return currentWordNode.Steps
		}

		for i := 0; i < len(currentWord); i++ {
			for j := 0; j < len(words); j++ {

				newWord := currentWord[:i] + string(words[j]) + currentWord[i+1:]

				for dictionaryWordIndex := range dictionaty {

					if dictionaty[dictionaryWordIndex] == newWord {
						dictionaty = append(dictionaty[:dictionaryWordIndex], dictionaty[dictionaryWordIndex+1:]...)
						queue = append([]Node{{
							Word:  newWord,
							Steps: currentWordNode.Steps + 1,
						}}, queue...)
						break
					}
				}
			}
		}
	}

	return 0
}
