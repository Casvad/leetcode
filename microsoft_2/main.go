package main

import "fmt"

// brute force solution
func Solution(S string) string {
	ans := ""
	process := false
	for i := 0; i < len(S); i++ {
		if i < len(S)-1 && (S[i] == 'A' && S[i+1] == 'B' || S[i] == 'B' && S[i+1] == 'A' || S[i] == 'C' && S[i+1] == 'D' || S[i] == 'D' && S[i+1] == 'C') {
			process = true
			i++
		} else {
			ans += string(S[i])
		}
	}
	if !process {
		return S
	}
	return Solution(ans)
}

func main() {

	fmt.Printf("%s\n", Solution("ACBDACBD"))
}
